import { Link, IconButton, Popover, PopoverArrow, PopoverBody, PopoverCloseButton, PopoverContent, PopoverHeader, PopoverTrigger, Flex, Box, Avatar, Text, Button, HStack } from "@chakra-ui/react";
import { Link as ReactLink } from 'react-router-dom';

import { FiPlus } from "react-icons/fi";
import { useAuth } from "../../../hooks/useAuth";
export function Navbar() {
    const { user, onSigOut } = useAuth();
    return (
        <Flex
            mr="4"
            as="aside"
        >
            <HStack mr="20" spacing={4}>
                <Link
                    as={ReactLink}
                    to='/campaigns'
                    py="1"
                    transition=" color 0.2s"
                    _hover={{ textDecoration: 'none', color: 'blue.900' }}
                >Campanhas
                </Link>
                <Link
                    as={ReactLink}
                    to='/contacts'
                    py="1"
                    transition=" color 0.2s"
                    _hover={{ textDecoration: 'none', color: 'blue.900' }}
                >Contatos
                </Link>
            </HStack>


            <Flex align="center">
                <Popover
                    placement="bottom-end">
                    <PopoverTrigger >
                        <Box
                            _hover={{ cursor: 'pointer' }}
                        >
                            <Avatar
                                size="md"
                                name={user?.name}
                            />
                        </Box>
                    </PopoverTrigger>
                    <PopoverContent color="gray.800" w="200px">
                        <Button
                            _hover={{ bg: "blue.900", color: 'white' }}
                            transition="background 0.2s"
                            type="button" variant="unstyled">
                            <Link
                                as={ReactLink}
                                to='/profile'
                                _hover={{ textDecoration: 'none' }}
                            >
                                Profile
                            </Link>
                        </Button>
                        <Button
                            _hover={{ bg: "blue.900", color: 'white' }}
                            transition="background 0.2s"
                            type="button" onClick={onSigOut} variant="unstyled">Sair</Button>
                    </PopoverContent>
                </Popover>
            </Flex>
        </Flex>
    )
}