import { destroyCookie, parseCookies, setCookie } from "nookies";
import { createContext, ReactNode, useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router-dom";
import { login } from "../services/api";
import { api } from "../services/apiClient";

interface AuthProviderProps {
    children: ReactNode;
}

interface User {
    name: string;
    email: string;
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

export const onSigOut = async () => {
    destroyCookie(undefined, '@Subscriber.token');
    destroyCookie(undefined, '@Subscriber.refreshToken');
}


export function AuthProvider({ children }: AuthProviderProps) {
    const [user, setUser] = useState({} as User);
    const isAuthenticated = !!user.name;
    const broadcastAuth = useRef<BroadcastChannel>(null);
    const navigate = useNavigate();

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
            api.get('/me').then(response => {


            }).catch(() => {
                onSigOut();
            });
        }

    }, []);

    const onLogin = async ({ email, password }: { email: string, password: string }) => {
        const response = await login({ email, password });
        console.log(response);
        // @ts-ignore
        const { user, token, refreshToken } = response.data;

        setCookie(undefined, '@Subscriber.token', token, {
            maxAge: 60 * 60 * 24 * 30, // 30 days
            path: '/'
        });
        setCookie(undefined, '@Subscriber.refreshToken', refreshToken, {
            maxAge: 60 * 60 * 24 * 30, // 30 days
            path: '/'
        });
        // @ts-ignore
        api.defaults.headers['Authorization'] = `Bearer ${token}`
        setUser(user);
        navigate('/dashboard');

    }


    return (
        <AuthContext.Provider value={{ user, isAuthenticated, broadcastAuth, onLogin, onSigOut }}>
            {children}
        </AuthContext.Provider>
    )
}