import { Button, Flex, Stack, useToast, Link, HStack, Heading, Icon, Box } from "@chakra-ui/react";
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { Link as ReactLink, useNavigate } from 'react-router-dom';
import { Layout } from "../../components/templates/Layout";
import { Input } from "../../components/utils/Input";
import { api } from "../../services/apiClient";
import { BiArrowBack } from "react-icons/bi";


interface FormProps {
    name: string;
    email: string;
}
const validation = Yup.object().shape({
    name: Yup.string().required('Nome é obrigatório'),
    email: Yup.string().email().required('E-mail é obrigatório'),
});


export function CreateClient() {
    const { register, handleSubmit, reset, formState } = useForm({
        resolver: yupResolver(validation)
    });
    const { errors, isSubmitting } = formState;
    const toast = useToast();
    const navigate = useNavigate();

    const onHandleSubmit: SubmitHandler<FormProps | FieldValues> = async (values) => {
        console.log(values);
        api.post('clients', {
            name: values.name,
            email: values.email,
        }).then((response) => {
            if (response.status === 201) {
                toast({
                    description: 'Cliente adicionado com sucesso!',
                    status: 'success',
                    duration: 5000,
                    isClosable: true
                });
                navigate('/clients');

            }
        }).catch(err => console.log(err))


    }
    return (
        <Layout>
            <Flex
                w="100vw"
                maxW={1480}
                height="500px"
                flexDirection="column"
                justify="center"
                align="center">

                <Flex w="100%"
                    maxW={[350, 800]}
                    ml={["-10", ""]}
                    mb="5"
                    justify="space-between"
                    align="center">
                    <Heading fontSize="2xl">Adicionar Cliente</Heading>
                    <Link
                        as={ReactLink}
                        to="/clients"
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
                    <Flex
                        onSubmit={handleSubmit(onHandleSubmit)}
                        as="form"
                        flex="1"
                        flexDirection="column"
                        justify="space-between"
                    >
                        <Stack spacing="4">
                            <Input
                                {...register('name')}
                                type="text"
                                label='Nome'
                                error={errors.name}

                            />
                            <Input
                                {...register('email')}
                                type="email"
                                label='E-mail'
                                error={errors.email}
                            />

                        </Stack>
                        <Box>
                            <Button
                                type="submit"
                                isLoading={isSubmitting}
                                transition="filter 0.2s"
                                _hover={{ filter: "brightnss(0.9)" }}
                                bg="blue.900">Salvar
                            </Button>
                        </Box>


                    </Flex>
                </Flex>

            </Flex>
        </Layout>
    )
}