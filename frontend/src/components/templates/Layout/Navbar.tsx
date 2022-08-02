import { Link, IconButton, Popover, PopoverArrow, PopoverBody, PopoverCloseButton, PopoverContent, PopoverHeader, PopoverTrigger, Flex, Box, Avatar, Text, Button } from "@chakra-ui/react";
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
            <Popover
                placement="bottom-end"
            >
                <PopoverTrigger>
                    <IconButton
                        color="gray.50"
                        fontSize="xl"
                        variant="unstyled"
                        aria-label='action button'
                        icon={<FiPlus />}
                    />
                </PopoverTrigger>
                <PopoverContent
                    maxWidth={200}
                    //px="4"
                    py="1"
                    textAlign="center"
                    gap="2"
                    color="gray.900">
                    <Flex
                        flexDirection="column"
                        w="100%"
                    >
                        <Link
                            as={ReactLink}
                            to='/campaigns'
                            py="1"
                            transition="background 0.2s, color 0.2s"
                            _hover={{ textDecoration: 'none', background: "blue.900", color: 'gray.50' }}
                        >Campanhas
                        </Link>

                    </Flex>
                </PopoverContent>
            </Popover>


            <Flex align="center">
                <Popover
                    placement="bottom-end">
                    <PopoverTrigger>
                        <Avatar
                            size="md"
                            name={user?.name}
                        />
                    </PopoverTrigger>
                    <PopoverContent color="gray.800" w="200px">
                        <Button type="button" onClick={onSigOut} variant="unstyled">Sair</Button>
                    </PopoverContent>
                </Popover>
            </Flex>
        </Flex>
    )
}