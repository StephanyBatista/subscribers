import { Link, IconButton, Popover, PopoverArrow, PopoverBody, PopoverCloseButton, PopoverContent, PopoverHeader, PopoverTrigger, Flex, Box, Avatar, Text } from "@chakra-ui/react";
import { Link as ReactLink } from 'react-router-dom';

import { FiPlus } from "react-icons/fi";
export function Navbar() {
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
                <Avatar
                    size="md"
                    name="Felipe Almeida Batista"
                />
            </Flex>
        </Flex>
    )
}