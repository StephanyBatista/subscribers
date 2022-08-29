import { Link, Button, Flex, Heading, HStack, Stack, useDisclosure, useToast, Icon, FormControl, FormLabel, Textarea, Alert, AlertDescription, Spinner, Grid, GridItem, Text, Divider, InputGroup, InputRightElement, ListItem, List } from "@chakra-ui/react";
import { useCallback, useEffect, useState } from "react";
import { useNavigate, useParams, Link as ReactLink } from "react-router-dom";
import { Layout } from "../../components/templates/Layout";
import { Input } from "../../components/utils/Input";
import { api } from "../../services/apiClient";
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { BiArrowBack } from "react-icons/bi";
import { AiOutlinePaperClip } from "react-icons/ai";
import { useDropzone } from "react-dropzone";

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
    status: string | boolean;
    baseofSubscribers: number;
    totalRead: number;
    totalSent: number;
    attachmentURL?: string;

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
    const [attachment, setAttachment] = useState<File>();
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

                if (response.data) {
                    toast({
                        description: "Campanha Ativa",
                        status: 'success',
                        duration: 5000,
                        isClosable: true
                    });
                    handleUpdateState();
                }
            }).catch((err) => {
                console.log(err)
            }).finally(() => setIsSendEmail(false));
    }, [])

    const onDrop = useCallback((acceptedFiles: File[]) => {

        switch (acceptedFiles[0].type) {
            case 'application/pdf':
                break;

            case 'application/msword':
                break;

            default:
                toast({
                    title: "Formato de arquivo inválido!",
                    description: "Aceito somente arquivos, pdf ou .doc, docx",
                    status: 'error',
                    duration: 9000,
                    isClosable: true
                })
                return
        }


        if (acceptedFiles[0].size >= 10735049) {
            toast({

                description: "Tamanho excedido, maximo permitido por arquivo é 10MB.",
                status: 'error',
                duration: 9000,
                isClosable: true
            })

            return false;
        }
    }, []);


    const { acceptedFiles, getRootProps, getInputProps } = useDropzone({
        maxFiles: 1,
        onDrop,
    });

    const files = acceptedFiles.map(file => (
        //@ts-ignore
        <ListItem key={file.path}>
            {/*@ts-ignore */}
            {file.path} - {file.size} bytes
        </ListItem>
    ));


    useEffect(() => {
        const controller = new AbortController();
        try {
            api.get<CampaignData>(`campaigns/${campaignId}`, { signal: controller.signal })
                .then((response) => {
                    console.log(response)
                    setCampaign({
                        body: response.data.body,
                        subject: response.data.subject,
                        createdAt: response.data.createdAt,
                        createdBy: response.data.createdBy,
                        from: response.data.from,
                        id: response.data.id,
                        name: response.data.name,
                        totalSent: response.data.totalSent,
                        totalRead: response.data.totalRead,
                        baseofSubscribers: response.data.baseofSubscribers,
                        status: response.data.status,
                        attachmentURL: response.data.attachmentURL,
                    })
                }).catch(err => console.log(err))
                .finally(() => setIsLoading(false));
        } catch (error) {

        }

        return () => { controller.abort() };

    }, [updateState]);


    const handleAttachment = async (event: React.ChangeEvent<HTMLInputElement>) => {
        const input = event.target;
        if (!input.files?.length) {
            return;
        }
        const file = input.files[0];

        switch (file.type) {
            case 'application/pdf':
                break;

            case 'application/msword':
                break;

            default:
                toast({
                    title: "Formato de arquivo inválido!",
                    description: "Aceito somente arquivos, pdf ou .doc, docx",
                    status: 'error',
                    duration: 9000,
                    isClosable: true
                })
                return
        }
        const formData = new FormData();
        formData.append('file', file);
        formData.append('kind', 'campaign');
        formData.append('keyId', String(campaignId));

        api.post('/files', formData).then((response) => {
            console.log(response)
        }).catch(err => console.log(err))
            .finally()



    }


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
                    w="100%"
                    maxW={[350, 800]}
                    ml={["-10", ""]}
                    mb="5"
                    justify="space-between"
                    align="center">
                    <Grid templateColumns={["1fr", "1fr 1fr 1fr"]} w="100%" gap="4" >
                        <GridItem >
                            <Flex
                                bg="gray.800"
                                px="2"
                                py="2"
                                w="100%"
                                h="100px"
                                maxHeight={150}
                                align="center"
                                flexDirection="column"
                                justify="space-between">
                                <Text fontSize="small" fontWeight="semibold">Total de Clientes</Text>
                                <Divider />
                                <Text fontSize="3xl" fontWeight="semibold" color="green">{campaign?.baseofSubscribers}</Text>
                            </Flex>
                        </GridItem>
                        <GridItem >
                            <Flex
                                bg="gray.800"
                                px="2"
                                py="2"
                                w="100%"
                                h="100px"
                                maxHeight={150}
                                align="center"
                                flexDirection="column"
                                justify="space-between">
                                <Text fontSize="small" fontWeight="semibold">Total enviados</Text>
                                <Divider />
                                <Text fontSize="3xl" fontWeight="semibold" color="blue">{campaign?.totalSent}</Text>
                            </Flex>
                        </GridItem>
                        <GridItem >
                            <Flex
                                bg="gray.800"
                                px="2"
                                py="2"
                                w="100%"
                                h="100px"
                                maxHeight={150}
                                align="center"
                                flexDirection="column"
                                justify="space-between">
                                <Text fontSize="small" fontWeight="semibold">Totais abertos</Text>
                                <Divider />
                                <Text fontSize="3xl" fontWeight="semibold" color="yellow">{campaign?.totalRead}</Text>
                            </Flex>
                        </GridItem>
                    </Grid>
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
                        <Grid templateColumns={["1fr", "1fr 1fr"]} gap="4">
                            <GridItem>
                                <Input
                                    {...register('name')}
                                    isDisabled={campaign?.status === 'Draft' ? false : true}
                                    error={errors.name}
                                    type="text"
                                    label="Nome"
                                    defaultValue={campaign?.name}
                                />
                            </GridItem>
                            <GridItem>
                                <Input
                                    {...register('subject')}
                                    isDisabled={campaign?.status === 'Draft' ? false : true}
                                    error={errors.subject}
                                    type="text"
                                    label="Assunto"
                                    defaultValue={campaign?.subject}
                                />
                            </GridItem>

                        </Grid>
                        <Grid templateColumns={["1fr", "1fr 1fr"]} gap="4">
                            <GridItem>
                                <Input
                                    {...register('from')}
                                    isDisabled={campaign?.status === 'Draft' ? false : true}
                                    error={errors.from}
                                    type="email"
                                    label="De"
                                    defaultValue={campaign?.from}
                                />
                            </GridItem>
                            {campaign?.attachmentURL ? (
                                <GridItem>
                                    <FormControl>
                                        <FormLabel mb="5">Arquivo</FormLabel>
                                        <Link bg="transparent"
                                            borderWidth={1}
                                            borderColor="blue.900"
                                            px="4"
                                            py="2"
                                            _hover={{ filter: "brightness(0.9)", textDecoration: 'none', bg: "blue.900" }}
                                            fontWeight="bold"
                                            href={campaign.attachmentURL}
                                            target="_blank"
                                        >Baixar</Link>
                                    </FormControl>

                                </GridItem>
                            ) : (
                                <GridItem>
                                    <Flex as="section" flexDirection="column" >
                                        <Text fontWeight="600">Arquivo*:</Text>
                                        <Flex
                                            border="dashed"
                                            py="8"
                                            px="8"
                                            borderWidth={1}
                                            {...getRootProps({ className: 'dropzone' })}
                                            align="center"
                                            justify="center"
                                        >
                                            <input type="file" {...register('arquivo')} name="arquivo" {...getInputProps()} />
                                            <Text fontSize="12">Clique aqui para selecionar ou arraste o arquivo</Text>
                                        </Flex>
                                        {campaign?.attachmentURL && (
                                            < Link fontWeight="bold" href={campaign.attachmentURL} target="_blank" fontSize="small">Baixar</Link>
                                        )}

                                    </Flex>

                                    {/* <FormControl>
                                     <FormLabel>Anexo da campanha</FormLabel>
                                     <InputGroup
                                         alignItems="center" >
                                         <InputRightElement
                                             pointerEvents="none"
                                             children={<Icon as={AiOutlinePaperClip} fontSize="20" />}
                                         />
                                         <Input
                                             isDisabled={campaign?.status === 'Draft' ? false : true}
                                             type="file"
                                             name="file"
                                             accept="application/pdf,application/msword,.docx"
                                             onChange={handleAttachment}
                                         />
                                     </InputGroup>
                                 </FormControl> */}
                                </GridItem>
                            )}

                        </Grid>

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

                    <Flex mt="10" justify="space-between" >
                        <HStack >
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
                        <Text fontSize="small" fontWeight="semibold" color="gray.300">Status: {campaign.status === "Processing" && "Processando" || campaign.status === "Draft" && "Rascunho"}</Text>
                    </Flex>

                </Flex>
            </Flex>


        </Layout >
    )
}