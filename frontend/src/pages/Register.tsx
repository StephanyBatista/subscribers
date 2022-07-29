import { Box, Button, Flex, Grid, GridItem, Heading, Stack, Link, useToast } from "@chakra-ui/react";
import { Link as ReactLink, useNavigate } from 'react-router-dom';
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { Input } from "../components/utils/Input";
import { api } from "../services/apiClient";

interface FormProps {
    email: string;
    name: string;
    password: string;
    confirmed_password: string;
}

const validation = Yup.object().shape({
    email: Yup.string().email().required('E-mail é obrigatório'),
    name: Yup.string().required('Nome é obrigatório'),
    password: Yup.string()
        .min(8, 'A senha precisa ter ao menos 8 caracteres')
        .required('A senha é obrigatório'),
    confirmed_password: Yup.string().oneOf(
        [Yup.ref('password')],
        'Senha não confere',
    ),
})


export function Register() {
    const { register, handleSubmit, reset, formState } = useForm({
        resolver: yupResolver(validation)
    });
    const { errors, isSubmitting } = formState;
    const toast = useToast();
    const navigate = useNavigate();

    const onHandleSubmit: SubmitHandler<FormProps | FieldValues> = async (values) => {
        const formData = new FormData();
        formData.append('name', values.name);
        formData.append('email', values.email);
        formData.append('password', values.password);
        const response = await api.post('/users/', {
            name: values.name,
            email: values.email,
            password: values.password
        });
        if (response.status === 201) {
            toast({
                description: "Cadastro realizado com sucesso!",
                status: 'success',
                duration: 5000,
                isClosable: true
            });
            navigate('/');
        } else {
            throw new Error("Erro ao salvar informações");

        }
    }
    return (
        <Flex
            as="form"
            onSubmit={handleSubmit(onHandleSubmit)}
            w="100vw"
            maxWidth={1100}
            height="100vh"
            justify="center"
            align="center"
            mx="auto"
        >
            <Grid templateColumns={["1fr", "1fr 1fr"]} >
                <GridItem>
                    <Flex borderRadius={8} bg="gray.800" w="100%" maxWidth={480} px="6" py="10" flexDirection="column">
                        <Stack spacing={2} w="100%">
                            <Input
                                {...register('name')}
                                label="Nome"

                                type="text"
                                error={errors.name}
                            />
                            <Input
                                {...register('email')}
                                label="E-mail"

                                type="email"
                                error={errors.email}
                            />
                            <Input
                                {...register('password')}
                                label="Senha"

                                type="password"
                                error={errors.password}
                            />
                            <Input
                                {...register('confirmed_password')}
                                label="Confirmar senha"

                                type="password"
                                error={errors.confirmed_password}
                            />
                        </Stack>
                        <Flex justify="center" mt="8">
                            <Button isLoading={isSubmitting} transition="filter 0.2s" bg="blue.900" width={300} _hover={{ filter: "brightness(0.9)" }} type="submit">Cadastrar</Button>
                        </Flex>
                    </Flex>
                </GridItem>
                <GridItem>
                    <Flex height="100%" align="center" mt={["4", "-4"]}>
                        <Stack spacing={4} ml={["4", ""]}  >
                            <Box  >
                                <Heading as="span" >Subs</Heading>
                                <Heading color="blue.900" as="span" >criber</Heading>
                            </Box>
                            <Heading fontSize="2xl">Cadastra-se é rápido e fácil</Heading>
                            <Link transition="filter 0.2s" color="blue.900" fontWeight="semibold" _hover={{ textDecoration: 'none', filter: "brightness(0.9)" }} as={ReactLink} to="/">Voltar para o login</Link>
                        </Stack>
                    </Flex>
                </GridItem>
            </Grid>


        </Flex>
    )
}