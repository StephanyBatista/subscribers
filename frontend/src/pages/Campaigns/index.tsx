import { Box, Button, Flex, Heading, Link, Stack, Table, Tbody, Td, Th, Tr, Thead, Icon } from "@chakra-ui/react";
import { useCallback, useEffect, useState } from "react";
import { Link as ReactLink } from "react-router-dom";
import { Layout } from "../../components/templates/Layout";
import { api } from "../../services/apiClient";
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
    body: string;
    status: string;
    subject: string;
}

export function Campaigns() {
    const [campaigns, setCampaigns] = useState<CampaignsData[]>([]);

    useEffect(() => {
        const controller = new AbortController();
        try {
            api.get('/campaigns', { signal: controller.signal }).then((response) => {
                setCampaigns(response.data);
            }).catch((error) => {

            });
        } catch (error) {

        }
        return () => { controller.abort() };
    }, []);

    const data = campaigns?.map(campaign => {

        return {
            id: campaign.id,
            body: campaign.body,
            createdAt: Intl.DateTimeFormat('pt-BR').format(new Date(campaign.createdAt)),
            from: campaign.from,
            name: campaign.name,
            subject: campaign.subject,
            createdBy: campaign.createdBy,
            status: campaign.status === "Processing" && "Processando" || campaign.status === "Draft" && "Rascunho"
        }
    })


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
                                    <Th>De</Th>
                                    <Th w="20">Status</Th>
                                    <Th w="20"></Th>
                                </Tr>
                            </Thead>
                            <Tbody>
                                {data?.map(campaign => (
                                    <Tr key={campaign.id} >
                                        <Td>{campaign.id}</Td>
                                        <Td>{campaign.name}</Td>
                                        <Td>{campaign.from}</Td>
                                        <Td> {campaign.status}</Td>
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