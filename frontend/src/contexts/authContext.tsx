import { destroyCookie, parseCookies, setCookie } from "nookies";
import { createContext, ReactNode, useCallback, useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Loading } from "../components/templates/Layout/Loading";
import { api } from "../services/apiClient";

interface AuthProviderProps {
    children: ReactNode;
}

interface User {
    name: string;
    email: string;
    userId: string;
}

interface AuthContextProps {
    user: User;
    broadcastAuth: any;
    onLogin: ({ email, password }: { email: string, password: string }) => Promise<void>;
    onSigOut: () => Promise<void>;
    isAuthenticated: boolean;
}

export const AuthContext = createContext({} as AuthContextProps);

let authChannel: BroadcastChannel


export async function onSigOut() {

    destroyCookie(undefined, '@Subscriber.token');
    destroyCookie(undefined, '@Subscriber.refreshToken');
    window.location.replace('/')
}


export function AuthProvider({ children }: AuthProviderProps) {
    const [user, setUser] = useState({} as User);
    const isAuthenticated = !!user?.name;
    const [isGetUser, setIsGetUser] = useState(true);
    const broadcastAuth = useRef<BroadcastChannel>(null);
    const navigate = useNavigate();

    useEffect(() => {
        if (!user?.name) {
            setIsGetUser(false);
        }
    }, [user]);


    useEffect(() => {
        // @ts-ignore
        broadcastAuth.current = new BroadcastChannel('auth');
        broadcastAuth.current.onmessage = (message) => {
            switch (message.data) {
                case 'signOut':
                    onSigOut();
                    break;
                default:
                    break;
            }
        }
    }, [broadcastAuth]);

    useEffect(() => {

        const { '@Subscriber.token': token } = parseCookies();
        if (token) {
            setIsGetUser(true);
            onGetUserInfos();
        }

    }, []);

    const onGetUserInfos = useCallback(async () => {

        api.get('/users/info').then(response => {
            setUser(response.data);
            console.log(response.data)
        }).catch((errors) => {
            console.log(errors)
            // onSigOut();
        }).finally(() => {
            setIsGetUser(false);
        });
    }, [])

    async function onLogin({ email, password }: { email: string, password: string }) {
        const response = await api.post('/token', {
            email, password
        })

        // @ts-ignore
        const { user, token } = response.data;

        setCookie(undefined, '@Subscriber.token', token, {
            maxAge: 60 * 60 * 24 * 30, // 30 days
            path: '/',
            secure: true,
            sameSite: 'Strict'
        });
        // setCookie(undefined, '@Subscriber.refreshToken', refreshToken, {
        //     maxAge: 60 * 60 * 24 * 30, // 30 days
        //     path: '/'
        // });
        // @ts-ignore
        api.defaults.headers['Authorization'] = `Bearer ${token}`
        //setUser(user);
        await onGetUserInfos();
        navigate('/dashboard');

    }

    if (isGetUser) {
        return (
            <Loading />
        )
    }

    return (
        <AuthContext.Provider value={{ user, isAuthenticated, broadcastAuth, onLogin, onSigOut }}>
            {children}
        </AuthContext.Provider>
    )
}