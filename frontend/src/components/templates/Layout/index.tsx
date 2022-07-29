import { Flex, Text } from "@chakra-ui/react";
import { ReactNode } from "react"
import { Content } from "./Content";
import { Header } from "./Header";

interface LayoutProps {
    children: ReactNode;
}

export function Layout({ children }: LayoutProps) {
    return (
        <>
            <Flex flexDirection="column" h="100vh" justify="space-between" m="0 auto" >
                <Header />
                <Flex py="4" px="4" flex="1" height="100%">
                    <Content>
                        {children}
                    </Content>
                </Flex>
                <Flex bg="gray.800" w="100%" h="40px" justify="center" align="center">
                    <Text color="gray.300" fontSize="small">Subscriber company {new Intl.DateTimeFormat('pt-BR', { year: 'numeric' }).format(new Date())}</Text>
                </Flex>
            </Flex>
        </>
    )
}