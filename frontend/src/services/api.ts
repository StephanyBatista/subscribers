import axios,{AxiosError} from 'axios';
import {  parseCookies,setCookie } from "nookies";
import { onSigOut } from '../contexts/authContext';



let isRefreshing = false;
let faildedRequestsQueue:any = []; 



// export const api = axios.create({
//     baseURL: import.meta.env.VITE_API_BACKEND,
//     headers:{
//         Authorization: `Bearer ${token}`
//     }
// })

export function setupAPIClient(ctx  = undefined){
    let cookies = parseCookies(ctx);

     const api = axios.create({
        baseURL:import.meta.env.VITE_API_BACKEND,
        headers:{
            Authorization: `Bearer ${cookies['@Subscriber.token']}`
        }
    });
    
    api.interceptors.response.use(response =>{
        return response;
    },(error:AxiosError) =>{
        if(error.response?.status === 401){
            // @ts-ignore
            if(error.response.data?.code === 'token.expired'){
                //renovar o token
                cookies = parseCookies(ctx);
    
                const {'@Subscriber.refreshToken': refreshToken} = cookies;
                const originalConfig = error.config
    
                if(!isRefreshing){
                    isRefreshing = true;
    
                    api.post('/refresh',{
                        refreshToken,
                    }).then(response=>{
                        const {token} = response.data;
        
                         setCookie(ctx, '@Subscriber.token', token, {
                    maxAge: 60 * 60 * 24 * 30, // 30 days
                    path: '/'
                });
        
                setCookie(ctx, '@Subscriber.refreshToken', response.data.refreshToken, {
                    maxAge: 60 * 60 * 24 * 30, // 30 days
                    path: '/'
                });
                 // @ts-ignore
                api.defaults.headers['Authorization'] = `Bearer ${token}`;
     // @ts-ignore
                    faildedRequestsQueue.forEach(request => request.onSuccess(token))
                    faildedRequestsQueue = [];
    
                    }).catch(err =>{
                         // @ts-ignore
                        faildedRequestsQueue.forEach(request => request.onFailure(err))
                        faildedRequestsQueue = [];
                        onSigOut();
                       
    
                    }).finally(() =>{
                        isRefreshing = false;
                    });
                }
    
    
                return new Promise((resolve,reject)=>{
                    faildedRequestsQueue.push({
                        onSuccess: (token:string) => {
                             // @ts-ignore
                            originalConfig.headers['Authorization'] = `Bearer ${token}`
    
                            resolve(api(originalConfig));
                        },
                        onFailure: (error:AxiosError) => {
                            reject(error);
                        }
                    })
                })
                
            }else{
                //desloga usuário
                onSigOut();
                
            }
        }
        return Promise.reject(error);
        
    })
    return api;
}


 export async function login({ email, password }: { email: string, password: string }){
    const login = {
        email: 'felipeab10@hotmail.com',
        password:'12345678'
    }
   
    return new Promise((resolve,reject)=>{
        setTimeout(()=>{
            if(email === login.email && password === login.password){
                const data = {
                   user:{
                    name:'Felipe Almeida Batista',
                    email:'felipeab10@hotmail.com',
                   },
                   refreshToken:'ADADSDAKKJKADII@',
                    token:'DASDA@#!@$#6535654645ddsfsf'
                }
                resolve({data:data});
            }else{
                reject('Usuário e senha inválidos!')
            }
        },750)
    })
 }