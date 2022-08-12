import { Box, Button, Flex, Grid, GridItem, Heading, Stack, Text, Link, HStack, useToast } from "@chakra-ui/react";
import { Link as ReactLink, useNavigate } from 'react-router-dom';
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { Input } from "../components/utils/Input";
import { useAuth } from "../hooks/useAuth";
import { useEffect } from "react";

const validation = Yup.object().shape({
    email: Yup.string().email().required('E-mail é obrigatório'),
    password: Yup.string()
        .required('A senha é obrigatório'),
})
interface FormProps {
    email: string;
    password: string;
}
export function SingIn() {
    const { onLogin } = useAuth();
    const { register, handleSubmit, reset, formState } = useForm({
        resolver: yupResolver(validation)
    });
    const navigate = useNavigate();
    const { errors, isSubmitting } = formState;
    const { isAuthenticated } = useAuth();
    const toast = useToast();


    useEffect(() => {
        if (isAuthenticated) {
            navigate('/dashboard')
        }
    }, [isAuthenticated])


    const onHandleSubmit: SubmitHandler<FormProps | FieldValues> = async (values) => {
        const response = await onLogin({ email: values.email, password: values.password }).catch((errors) => {
            if (errors.response?.status === 403) {
                toast({
                    description: 'Usuário ou senha estão inválidos',
                    status: 'error',
                    duration: 9000,
                    isClosable: true
                })
            }
        });

    }

    return (
        <Flex
            as="form"
            onSubmit={handleSubmit(onHandleSubmit)}
            w="100vw"
            maxWidth={1100}
            justify="center"
            align="center"
            mx="auto"
            height="100vh"
        >
            <Grid templateColumns={["1fr", "1fr 1fr"]} >
                <GridItem>
                    <Flex height="100%" align='center'>
                        <Stack spacing={8} ml={["4", ""]}  >
                            <Box  >
                                <Heading as="span" >Subs</Heading>
                                <Heading color="blue.900" as="span" >criber</Heading>
                            </Box>
                            <Heading>Faça seu login na plataforma</Heading>
                        </Stack>
                    </Flex>
                </GridItem>
                <GridItem w="100%" maxWidth={480}>
                    <Flex bg="gray.800" borderRadius={8} m={["2", ""]} py="10" flexDirection="column">
                        <Stack spacing="8">
                            <Stack spacing={4} w="22rem" mx="auto" >
                                <Input
                                    _placeholder={{ color: 'gray.400', fontSize: '14px' }}
                                    {...register('email')}
                                    type="email"
                                    label="E-mail"
                                    error={errors.email}
                                />
                                <Input
                                    _placeholder={{ color: 'gray.400', fontSize: '14px' }}
                                    {...register('password')}
                                    type="password"
                                    label="Senha"
                                    error={errors.password}
                                />
                            </Stack>
                            <Flex  >

                                <Stack spacing={4} w="100%" justify="center" align="center">
                                    <Button w="22rem" bg="blue.900" transition="filter 0.2s" _hover={{ filter: "brightness(0.9)" }} type="submit">Entrar</Button>
                                    <HStack >
                                        <Text fontWeight="light">Não tem uma conta?</Text>
                                        <Link transition="filter 0.2s" color="blue.900" fontSize="lg" fontWeight="bold" _hover={{ textDecoration: 'none', filter: "brightness(0.9)" }} as={ReactLink} to="/register">Registre-se</Link>
                                    </HStack>
                                </Stack>
                            </Flex>
                        </Stack>
                    </Flex>
                </GridItem>
            </Grid>

        </Flex>
    );
}

