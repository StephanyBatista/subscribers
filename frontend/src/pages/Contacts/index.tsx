import { Box, Button, Flex, Heading, Icon, Link, Spinner, Stack, Table, Tbody, Td, Th, Thead, Tr } from "@chakra-ui/react";
import { Link as ReactLink } from "react-router-dom";
import { BiPencil } from "react-icons/bi";
import { useCallback, useEffect, useState } from "react";
import { Layout } from "../../components/templates/Layout";
import { api } from "../../services/apiClient";


interface ContactData {
    id: string;
    name: string;
    email: string;
}

export function Contacts() {
    const [contacts, setContacts] = useState<ContactData[]>([]);
    const [isLoading, setIsLoading] = useState(true);

    const getAllContacts = useCallback(async () => {
        api.get('/clients').then((response) => {
            setContacts(response.data)
        }).catch(err => console.log(err)).finally(() => setIsLoading(false));

    }, []);

    useEffect(() => {
        getAllContacts();
    }, [])

    return (
        <Layout>
            <Flex justify="space-between" mb="8" align="center">
                <Heading>Contatos</Heading>
                <Link
                    _hover={{ textDecoration: 'none' }}
                    as={ReactLink}
                    to="/contacts/create"
                >
                    <Button
                        type="button"
                        transition="filter 0.2s"
                        bg="blue.900"
                        _hover={{ filter: "brightness(0.9)" }}
                    >Adicionar
                    </Button>
                </Link>
            </Flex>
            <Flex
                px={["2", "8"]}
                ml={["-6", ""]}
                py={["2", "8"]}
                h="100%"
                w="100vw"
                maxW={1480}
                justify="flex-start"
                mx="auto"
                bg="gray.800"
                borderRadius="8"
                flexDirection="column"
            >
                <Stack spacing={8}>
                    {isLoading ? (
                        <Flex w="100%" flex="1" justify="center" align="center" mx="auto">
                            <Spinner />
                        </Flex>
                    ) : (
                        <Box overflowY="auto">
                            <Table colorScheme='whiteAlpha'>
                                <Thead>
                                    <Tr>
                                        <Th w="16">#</Th>
                                        <Th >Nome</Th>
                                        <Th >E-mail</Th>
                                        <Th w="20"></Th>
                                    </Tr>
                                </Thead>
                                <Tbody>
                                    {contacts?.map(contact => (
                                        <Tr key={contact.id} >
                                            <Td>{contact.id}</Td>
                                            <Td>{contact.name.charAt(0).toUpperCase() + contact.name.slice(1)}</Td>
                                            <Td>{contact.email}</Td>
                                            <Td>
                                                <Link
                                                    _hover={{ textDecoration: 'none' }}
                                                    as={ReactLink}
                                                    to={`/contacts/edit/${contact.id}`}
                                                >
                                                    <Button
                                                        type="button"
                                                        transition="filter 0.2s"
                                                        bg="blue.900"
                                                        _hover={{ filter: "brightness(0.9)" }}
                                                    >
                                                        <Icon as={BiPencil} />
                                                    </Button>
                                                </Link>
                                            </Td>

                                        </Tr>
                                    ))}
                                </Tbody>
                            </Table>
                        </Box>
                    )}

                </Stack>
            </Flex>
        </Layout>
    )
}