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
                <Flex my="6" px="6" w="100%" maxWidth={1480} mx="auto">
                    <Flex flexDirection="column" minHeight="calc(100vh - 200px)">
                        <Content>
                            {children}
                        </Content>
                    </Flex>
                </Flex>
                <Flex bg="gray.800" w="100%" h="40px" justify="center" align="center">
                    <Text color="gray.300" fontSize="small">Subscriber company {new Intl.DateTimeFormat('pt-BR', { year: 'numeric' }).format(new Date())}</Text>
                </Flex>
            </Flex>
        </>
    )
}