import { Button, Flex, Stack, useToast, Link, HStack, Heading } from "@chakra-ui/react";
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { Link as ReactLink, useNavigate } from 'react-router-dom';
import { Layout } from "../../components/templates/Layout";
import { Input } from "../../components/utils/Input";
import { api } from "../../services/apiClient";


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
        console.log(values)

        const response = await api.post('clients', {
            name: values.name,
            email: values.email,
        });
        if (response.status === 201) {
            toast({
                description: 'Cliente adicionado com sucesso!',
                status: 'success',
                duration: 5000,
                isClosable: true
            });
            navigate('/clients');

        }
        console.log(response);
    }
    return (
        <Layout>
            <Flex justify="space-between" mb="8" align="center">
                <Heading fontSize="2xl">Adicionar Clientes</Heading>
                <Link
                    _hover={{ textDecoration: 'none' }}
                    as={ReactLink}
                    to="/clients/create"
                >

                </Link>
            </Flex>
            <Flex
                px={["2", "8"]}
                ml={["-6", ""]}
                py={["2", "8"]}
                h="100%"
                w="100vw"
                maxW={600}
                maxH={400}
                justify="space-between"
                mx="auto"
                bg="gray.800"
                borderRadius="8"
                flexDirection="column"
                as="form"
                onSubmit={handleSubmit(onHandleSubmit)}
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
                <HStack>
                    <Link
                        py="2"
                        px="4"
                        borderRadius="6"
                        transition="filter 0.2s"
                        _hover={{ filter: "brightness(0.9)" }}
                        bg="gray.600"
                        as={ReactLink}
                        mr={3}
                        to="/clients"
                    >
                        Voltar
                    </Link>
                    <Button
                        type="submit"
                        isLoading={isSubmitting}
                        transition="filter 0.2s"
                        _hover={{ filter: "brightnss(0.9)" }}
                        bg="blue.900">Salvar
                    </Button>
                </HStack>



            </Flex>
        </Layout>
    )
}