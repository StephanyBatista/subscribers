import { Link, IconButton, Popover, PopoverArrow, PopoverBody, PopoverCloseButton, PopoverContent, PopoverHeader, PopoverTrigger, Flex, Box, Avatar, Text, Button, HStack } from "@chakra-ui/react";
import { Link as ReactLink } from 'react-router-dom';

import { FiPlus } from "react-icons/fi";
import { useAuth } from "../../../hooks/useAuth";
export function Navbar() {
    const { user, onSigOut, isAuthenticated } = useAuth();

    return (
        <Flex
            mr="4"
            as="aside"
        >
            {isAuthenticated ? (
                <>
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
                                <Button type="button" onClick={onSigOut} variant="unstyled">Sair</Button>
                            </PopoverContent>
                        </Popover>
                    </Flex>
                </>
            ) : (
                <HStack mr="20" spacing={4}>
                    <Link
                        as={ReactLink}
                        to='/singin'
                        py="1"
                        transition=" color 0.2s"
                        _hover={{ textDecoration: 'none', color: 'blue.900' }}
                    >Entrar
                    </Link>
                    <Link
                        as={ReactLink}
                        to='/register'
                        py="1"
                        transition=" color 0.2s"
                        _hover={{ textDecoration: 'none', color: 'blue.900' }}
                    >Registrar
                    </Link>
                </HStack>
            )}

        </Flex>
    )
}