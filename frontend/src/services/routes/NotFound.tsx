import { Flex, Heading, HStack, Link, Text } from "@chakra-ui/react";
import { Link as ReactLink } from "react-router-dom";

export function NotFound() {
    return (
        <Flex
            w="100vw"
            justify="center"
            align="center"
            height="100vh"
            maxWidth={1100}
            mx="auto"
            flexDirection="column"
        >
            <Flex align="center" px="2" py="2">
                <HStack>
                    <Heading>404</Heading>
                    <Text>Página não encontrada</Text>
                </HStack>
            </Flex>
            <Link
                transition="filter 0.2s"
                color="blue.900"
                fontWeight="semibold"
                _hover={{ textDecoration: 'none', filter: "brightness(0.9)" }}
                as={ReactLink}
                to="/"
            >Voltar</Link>
        </Flex>
    )
}