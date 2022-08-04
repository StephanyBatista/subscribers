import { ChakraProvider, Flex } from "@chakra-ui/react";
import { useEffect } from "react";
import { theme } from "../../../styles/theme";
import { Logo } from "./Logo";

export function Loading() {

    useEffect(() => {
        const timer = setTimeout(() => {

        }, 5000)

        return () => clearTimeout(timer);
    })
    return (
        <ChakraProvider theme={theme} resetCSS>
            <Flex
                justify="center"
                align="center"
                w="100%"
                h="100vh"
                bg="gray.900"
                flex="1"
            >
                <Logo />
            </Flex>
        </ChakraProvider>

    )
}