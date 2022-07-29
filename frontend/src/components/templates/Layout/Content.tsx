import { Flex } from "@chakra-ui/react";
import { ReactNode } from "react";
interface ContentProps {
    children: ReactNode;
}
export function Content({ children }: ContentProps) {
    return (
        <>
            {children}
        </>
    )
}