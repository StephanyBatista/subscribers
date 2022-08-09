import { Box, Link, Button, Flex, Heading, HStack, List, ListItem, Stack, Tab, Table, TabList, TabPanel, TabPanels, Tabs, Th, Thead, Tr, useDisclosure, useToast, Icon, FormControl, FormLabel, Textarea, Alert, AlertDescription, Spinner } from "@chakra-ui/react";
import { useCallback, useEffect, useState } from "react";
import { useNavigate, useParams, Link as ReactLink } from "react-router-dom";
import { Layout } from "../../components/templates/Layout";
import { Input } from "../../components/utils/Input";
import { api } from "../../services/apiClient";
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { ClientModal } from "../../components/campaigns/ClientModal";
import { BiArrowBack } from "react-icons/bi";


interface FormProps {
    name: string;
    description: string;
    active: boolean;
}

interface CreatedBy {
    Id: string;
    Name: string;
}

interface CampaignData {
    body: string;
    subject: string;
    createdAt: string;
    createdBy: CreatedBy;
    from: string;
    id: string;
    name: string;
    status: string;

}

const validation = Yup.object().shape({
    name: Yup.string().required('Nome é obrigatório'),
    description: Yup.string().required('A descrição é obrigatório'),
    //active: Yup.boolean()
    //.required("The terms and conditions must be accepted.s")
    //  .oneOf([true], "The terms and conditions must be accepted.")
})

export function Edit() {
    const [campaign, setCampaign] = useState({} as CampaignData);
    const [isLoading, setIsLoading] = useState(true);
    const [isSendEmail, setIsSendEmail] = useState(false);
    const [updateState, setUpdateState] = useState(false);
    const { campaignId } = useParams();

    const { register, handleSubmit, reset, formState } = useForm({
        resolver: yupResolver(validation)
    });
    const { errors, isSubmitting } = formState;
    const toast = useToast();
    const navigate = useNavigate();
    const { isOpen, onOpen, onClose } = useDisclosure()


    const handleUpdateState = () => {
        setUpdateState(!updateState);
    }
    const onHandleSubmit: SubmitHandler<FormProps | FieldValues> = async (values) => {

    }

    const handleSendEmails = useCallback(async () => {
        setIsSendEmail(true);
        api.post(`campaigns/${campaignId}/send`)
            .then((response) => {
                console.log(response);
            }).catch((err) => {
                console.log(err)
            }).finally(() => setIsSendEmail(false));
    }, [])

    const getCampaign = useCallback(async () => {
        api.get(`campaigns/${campaignId}`)
            .then((response) => {
                setCampaign({
                    body: response.data.body,
                    subject: response.data.subject,
                    createdAt: response.data.createdAt,
                    createdBy: response.data.createdBy,
                    from: response.data.from,
                    id: response.data.id,
                    name: response.data.name,
                    status: response.data.status,
                })
            }).catch(err => console.log(err))
            .finally(() => setIsLoading(false));
        // setCampaign(response.data);

    }, []);

    useEffect(() => {
        getCampaign();
    }, [updateState]);

    if (isLoading) {
        return (
            <Layout>
                <Flex
                    w="100vw"
                    maxW={1480}
                    flex="1"
                    justify="center"
                    align="center"
                >
                    <Spinner />
                </Flex>
            </Layout>
        )
    }
    return (
        <Layout>

            <Flex
                w="100vw"
                maxW={1480}
                flex="1"
                flexDirection="column"
                justify="center"
                align="center"
            >
                <Flex
                    w="100%"
                    maxW={[350, 800]}
                    ml={["-10", ""]}
                    mb="5"
                    justify="space-between"
                    align="center">
                    <Heading fontSize="xl">Gerenciar campanha</Heading>
                    <Link
                        as={ReactLink}
                        to="/campaigns"
                    >
                        <Icon as={BiArrowBack} fontSize="2xl" />
                    </Link>
                </Flex>
                <Flex
                    px={["2", "8"]}
                    ml={["-10", ""]}
                    py={["4", "8"]}
                    h="100%"
                    w="100%"
                    maxW={[350, 800]}
                    justify="space-between"
                    // mx="auto"
                    bg="gray.800"
                    borderRadius="8"
                    flexDirection="column"
                >
                    <Stack
                        as="form"
                        spacing={3}>

                        <Input
                            {...register('name')}
                            isDisabled={campaign?.status === 'Draft' ? false : true}
                            error={errors.name}
                            type="text"
                            label="Nome"
                            defaultValue={campaign?.name}
                        />
                        <Input
                            {...register('from')}
                            isDisabled={campaign?.status === 'Draft' ? false : true}
                            error={errors.from}
                            type="email"
                            label="De"
                            defaultValue={campaign?.from}
                        />
                        <Input
                            {...register('subject')}
                            isDisabled={campaign?.status === 'Draft' ? false : true}
                            error={errors.subject}
                            type="text"
                            label="Assunto"
                            defaultValue={campaign?.subject}
                        />
                        <FormControl>
                            <FormLabel>Texto</FormLabel>
                            <Textarea
                                {...register('body')}
                                isDisabled={campaign?.status === 'Draft' ? false : true}
                                resize="none"
                                bg="gray.950"
                                border="none"
                                defaultValue={campaign?.body}
                            />
                            {errors.body && (
                                <Alert bg="transparent" color="red.600" fontSize="0.875rem">
                                    <AlertDescription>
                                        {/* @ts-ignore */}
                                        {errors.body?.message}
                                    </AlertDescription>
                                </Alert>

                            )}
                        </FormControl>

                    </Stack>

                    <HStack mt="10">
                        <Button
                            disabled={campaign?.status === 'Draft' ? false : true}
                            type="submit"
                            transition="filter 0.2s"
                            _hover={{ filter: "brightness(0.9)" }}
                            bg="blue.900"
                        >Atualizar
                        </Button>
                        <Button
                            disabled={campaign?.status === 'Draft' ? false : true}
                            type="submit"
                            transition="filter 0.2s"
                            _hover={{ filter: "brightness(0.9)" }}
                            bg="green.700"
                            onClick={handleSendEmails}
                        >Disparar E-mails
                        </Button>
                    </HStack>

                </Flex>
            </Flex>


        </Layout>
    )
}