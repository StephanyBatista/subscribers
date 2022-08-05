import { Box, Link, Button, Flex, Heading, HStack, List, ListItem, Stack, Tab, Table, TabList, TabPanel, TabPanels, Tabs, Th, Thead, Tr, useDisclosure, useToast, Icon, Tbody, Td } from "@chakra-ui/react";
import { useCallback, useEffect, useState } from "react";
import { useNavigate, useParams, Link as ReactLink } from "react-router-dom";
import { Layout } from "../../components/templates/Layout";
import { Input } from "../../components/utils/Input";
import { api } from "../../services/apiClient";
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { ClientModal } from "../../components/campaigns/ClientModal";
import { BiArrowBack, BiPencil } from "react-icons/bi";


interface FormProps {
    name: string;
    description: string;
    active: boolean;
}

interface CreatedBy {
    Id: string;
    Name: string;
}

interface ClientData {
    active: boolean;
    createdAt: string;
    createdBy: CreatedBy;
    email: string;
    id: string;
    name: string;
}

interface CampaignData {
    active: boolean;
    createdAt: string;
    createdBy: CreatedBy;
    description: string;
    id: string;
    name: string;
    clients: ClientData[];

}

const validation = Yup.object().shape({
    name: Yup.string().required('Nome é obrigatório'),
    description: Yup.string().required('A descrição é obrigatório'),
    //active: Yup.boolean()
    //.required("The terms and conditions must be accepted.s")
    //  .oneOf([true], "The terms and conditions must be accepted.")
})

export function Edit() {
    const [campaign, SetCampaign] = useState({} as CampaignData);
    const [client, setClient] = useState({} as ClientData);
    const [updateState, setUpdateState] = useState(false);
    const { campaignId } = useParams();
    const { register, handleSubmit, reset, formState } = useForm({
        resolver: yupResolver(validation)
    });
    const { errors, isSubmitting } = formState;
    const toast = useToast();
    const navigate = useNavigate();
    const { isOpen, onOpen, onClose } = useDisclosure()

    const handleResetClient = () => {
        setClient({} as ClientData);
        console.log(client)
    }
    const handleUpdateState = () => {
        setUpdateState(!updateState);
    }
    const onHandleSubmit: SubmitHandler<FormProps | FieldValues> = async (values) => {

    }

    const getCampaign = useCallback(async () => {
        const response = await api.get(`campaigns/${campaignId}`);
        SetCampaign(response.data);
        console.log(response);
    }, []);

    useEffect(() => {
        getCampaign();
    }, [updateState]);


    return (
        <Layout>
            <Flex
                px={["2", "8"]}
                ml={["-6", ""]}
                py={["2", "8"]}
                h="100%"
                w="100vw"
                maxW={1480}
                justify="space-between"
                mx="auto"
                bg="gray.800"
                borderRadius="8"
                flexDirection="column"
            >

                {/* <HStack>
                    <Stack
                        as="form"
                        spacing={3}>
                        <Input
                            name=""
                            type="text"
                            label="Nome"
                        />
                        <Input
                            name=""
                            type="text"
                            label="Descrição"
                        />
                        <Box mt="10">
                            <Button
                                type="submit"
                                transition="filter 0.2s"
                                _hover={{ filter: "brightness(0.9)" }}
                                bg="blue.900"
                            >Atualizar
                            </Button>
                        </Box>
                    </Stack>
                    <List >
                        <ListItem><strong>Criado por: </strong>{campaign.CreatedBy.Name}</ListItem>
                        <ListItem><strong>as: </strong>{campaign.CreatedBy.Name}</ListItem>
                    </List>
                </HStack> */}

                <Flex flexDirection="column">
                    <Flex justify="space-between" align="center" mb="4">
                        <Heading fontSize="2xl">Clientes</Heading>
                        <HStack spacing={8} align="center">
                            <Link as={ReactLink} to="/campaigns">
                                <Icon as={BiArrowBack} fontSize={22} />
                            </Link>
                            <Button onClick={onOpen} colorScheme="green">Adicionar cliente</Button>
                        </HStack>
                    </Flex>
                    <Table colorScheme="whiteAlpha">
                        <Thead>
                            <Tr>
                                <Th w="8">#</Th>
                                <Th >Nome</Th>
                                <Th>E-mail</Th>
                                <Th w="8"></Th>
                            </Tr>
                        </Thead>
                        <Tbody>
                            {campaign.clients?.map(client => (
                                <Tr key={client.id}>
                                    <Td>{client.id}</Td>
                                    <Td>{client.name}</Td>
                                    <Td>{client.email}</Td>
                                    <Td>
                                        <Button
                                            type="button"
                                            onClick={() => { onOpen(); setClient(client); }}
                                            transition="filter 0.2s"
                                            bg="blue.900"
                                            _hover={{ filter: "brightness(0.9)" }}
                                        >
                                            <Icon as={BiPencil} />
                                        </Button>
                                    </Td>
                                </Tr>
                            ))}
                        </Tbody>
                    </Table>
                </Flex>
                {/* <Tabs variant="unstyled">
                    <TabList>
                        <Tab transition="color 0.2s, fontWeight 0.2s" _selected={{ color: 'blue.900', bg: 'transparent', fontWeight: 'bold' }}>Campanha</Tab>
                        <Tab transition="color 0.2s, fontWeight 0.2s" _selected={{ color: 'blue.900', bg: 'transparent', fontWeight: 'bold' }}>Clientes</Tab>
                    </TabList>

                    <TabPanels>
                        <TabPanel>
                            <p>one!</p>
                        </TabPanel>
                        <TabPanel>
                            <p>two!</p>
                        </TabPanel>
                        <TabPanel>
                            <p>three!</p>
                        </TabPanel>
                    </TabPanels>
                </Tabs> */}
            </Flex>
            <ClientModal
                isOpen={isOpen}
                isUpdating={!!client.id}
                client={client}
                campaignId={campaignId}
                onClose={() => { onClose(); handleResetClient() }}
                onUpdateState={handleUpdateState}
            />
        </Layout>
    )
}