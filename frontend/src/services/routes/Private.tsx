import { useToast } from "@chakra-ui/react";
import { ReactElement, ReactNode, useEffect } from "react"
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../hooks/useAuth";
import { Forbidden } from "./Forbidden";

interface PrivateProps {
    children: ReactNode;
}
export function Private({ children }: PrivateProps) {
    const { isAuthenticated } = useAuth();
    const toast = useToast();
    const navigate = useNavigate();
    console.log('isAuthenticated ', isAuthenticated);

    useEffect(() => {
        if (!isAuthenticated) {
            navigate('/singin');
        }
    }, [isAuthenticated])

    if (isAuthenticated) {
        return (
            <>
                {children}
            </>
        )
    }
    return null;



}