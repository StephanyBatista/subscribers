import { Box, Button, Flex, Grid, GridItem, Heading, Link, List, ListItem, Stack, Table, Tbody, Td, Th, Image, Text, Tr, Thead, Icon } from "@chakra-ui/react";
import { useCallback, useEffect, useState } from "react";
import { Link as ReactLink } from "react-router-dom";
import { Layout } from "../../components/templates/Layout";
import { api } from "../../services/apiClient";
import Test from "../../assets/test.jpeg"
import { FiPenTool } from "react-icons/fi";
import { BiPencil } from "react-icons/bi";


interface UserCreated {
    Id: string;
    Name: string;
}

interface CampaignsData {
    createdAt: string;
    createdBy: UserCreated;
    from: string;
    id: string;
    name: string;
}

export function Campaigns() {
    const [campaigns, setCampaigns] = useState<CampaignsData[]>([]);

    const getAllCampaigns = useCallback(async () => {
        const response = await api.get('/campaigns').then((response) => {
            console.log(response)
            setCampaigns(response.data);
        }).catch((error) => {
            console.log('rttot ', error)
        });


    }, []);

    useEffect(() => {
        getAllCampaigns();
    }, []);

    // async function test() {
    //     fetch('https://api.unsplash.com/photos', {
    //         method: 'get',
    //         headers: new Headers({
    //             'Authorization': `Client-ID 46FJgSYfVO3QivzUqsriSneBlY3Osq5N4TMXNpGumiE`,

    //         }),

    //     }).then(data => data.json())
    //         .then((resp) => {
    //             console.log(resp)
    //         })
    // }
    // useEffect(() => {
    //     test();
    // }, [])

    return (
        <Layout>
            <Flex justify="space-between" mb="8" align="center">
                <Heading>Minhas Campanhas</Heading>
                <Link
                    _hover={{ textDecoration: 'none' }}
                    as={ReactLink}
                    to="/campaigns/create"
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
            {/* <Grid templateColumns={["1fr", "repeat(4,1fr)"]} gap={8}>
                {campaigns.map(campaign => (
                    <GridItem key={campaign.ID}>
                        {campaign.ID ? (
                            <Image height="150px" w="100%" objectFit="cover" src={Test} />
                        ) : (
                            <Box w="100%" bg="red" height="150px"></Box>
                        )}
                        <Flex flexDirection="column" w="100%" bg="gray.800" border="none" borderBottomRadius={8} px="8" py="6">
                            <List spacing={3}>
                                <ListItem>
                                    <Heading fontSize="xl" as="span">Campanha:  </Heading>
                                    <Heading fontSize="md" as="span">{campaign.Name}</Heading>
                                </ListItem>
                                <ListItem>
                                    {campaign.Active ? (
                                        <Text px="1" w="70px" textAlign="center" borderRadius="4" py="1" color="gray.50" fontWeight="bold" bg="green.700">Ativo</Text>
                                    ) : (
                                        <Text px="1" w="70px" textAlign="center" borderRadius="4" py="1" color="gray.50" fontWeight="bold" bg="red.200">Desativado</Text>
                                    )}
                                </ListItem>
                            </List>
                        </Flex>


                    </GridItem>
                ))}
            </Grid> */}

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
                    {/* <Flex justify="space-between">
                        <Heading fontSize="2xl">Campanhas</Heading>
                        <Link
                            _hover={{ textDecoration: 'none' }}
                            as={ReactLink}
                            to="/campaigns/create"
                        >
                            <Button
                                type="button"
                                transition="filter 0.2s"
                                bg="blue.900"
                                _hover={{ filter: "brightness(0.9)" }}
                            >Adicionar
                            </Button>
                        </Link>
                    </Flex> */}
                    <Box overflowY="auto">
                        <Table colorScheme='whiteAlpha'>
                            <Thead>
                                <Tr>
                                    <Th w="16">#</Th>
                                    <Th >Nome</Th>
                                    <Th>De</Th>
                                    <Th w="20"></Th>
                                </Tr>
                            </Thead>
                            <Tbody>
                                {campaigns.map(campaign => (
                                    <Tr key={campaign.id} >
                                        <Td>{campaign.id}</Td>
                                        <Td>{campaign.name}</Td>
                                        <Td>{campaign.from}</Td>
                                        <Td>
                                            <Link
                                                _hover={{ textDecoration: 'none' }}
                                                as={ReactLink}
                                                to={`/campaigns/edit/${campaign.id}`}
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