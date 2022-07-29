import { Flex } from "@chakra-ui/react";
import { Logo } from "./Logo";
import { Navbar } from "./Navbar";

export function Header() {
    return (
        <Flex
            w="100vw"
            h="60px"
            px="4"
            align="center"
            bg="gray.800"
            justify="space-between"
        >
            <Logo />
            <Navbar />
        </Flex>
    )
}