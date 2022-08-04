import { Box, Button, Flex, Heading, Link, Stack, Table, Tbody, Td, Th, Thead, Tr } from "@chakra-ui/react";
import { useCallback, useEffect, useState } from "react";
import { Link as ReactLink } from "react-router-dom";
import { Layout } from "../../components/templates/Layout";
import { api } from "../../services/apiClient";

interface UserCreated {
    Id: string;
    Name: string;
}

interface CampaignsData {
    Active: boolean;
    CreatedAt: string;
    CreatedBy: UserCreated;
    Description: string;
    ID: string;
    Name: string;
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
    return (
        <Layout>
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
                    <Flex justify="space-between">
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
                    </Flex>
                    <Box overflowY="auto">
                        <Table colorScheme='whiteAlpha'>
                            <Thead>
                                <Tr>
                                    <Th w="16">#</Th>
                                    <Th >Nome</Th>
                                    <Th w="8">Status</Th>
                                </Tr>
                            </Thead>
                            <Tbody>
                                {campaigns.map(campaign => (
                                    <Tr key={campaign.ID}>
                                        <Td>{campaign.ID}</Td>
                                        <Td>{campaign.Name}</Td>
                                        {campaign.Active ? (
                                            <Td>Ativo</Td>
                                        ) : (
                                            <Td>Desativado</Td>
                                        )}

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