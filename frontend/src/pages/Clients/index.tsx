import { Box, Button, Flex, Heading, Icon, Link, Stack, Table, Tbody, Td, Th, Thead, Tr } from "@chakra-ui/react";
import { Link as ReactLink } from "react-router-dom";
import { BiPencil } from "react-icons/bi";
import { useCallback, useEffect, useState } from "react";
import { Layout } from "../../components/templates/Layout";


interface ClientsData {
    id: string;
    name: string;
    email: string;
}

export function Clients() {
    const [clients, setClients] = useState<ClientsData[]>([]);


    const getAllClientes = useCallback(async () => {
        fetch('http://localhost:3000/clients').then(data => data.json()).then((response) => {
            console.log(response)
            setClients(response)
        })
    }, []);

    useEffect(() => {
        getAllClientes();
    }, [])

    return (
        <Layout>
            <Flex justify="space-between" mb="8" align="center">
                <Heading>Clientes</Heading>
                <Link
                    _hover={{ textDecoration: 'none' }}
                    as={ReactLink}
                    to="/clients/create"
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

                    <Box overflowY="auto">
                        <Table colorScheme='whiteAlpha'>
                            <Thead>
                                <Tr>
                                    <Th w="16">#</Th>
                                    <Th >Nome</Th>
                                    <Th w="20"></Th>
                                </Tr>
                            </Thead>
                            <Tbody>
                                {clients.map(client => (
                                    <Tr key={client.id} >
                                        <Td>{client.id}</Td>
                                        <Td>{client.name}</Td>

                                        <Td>
                                            <Link
                                                _hover={{ textDecoration: 'none' }}
                                                as={ReactLink}
                                                to={`/clients/edit/${client.id}`}
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
                </Stack>
            </Flex>
        </Layout>
    )
}