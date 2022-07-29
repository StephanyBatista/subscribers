import { Link, Box, Button, Flex, FormControl, FormLabel, Heading, IconButton, Stack, Switch, Icon, useToast, FormErrorMessage, Alert, AlertIcon, AlertDescription } from "@chakra-ui/react";
import { Link as ReactLink, useNavigate } from 'react-router-dom';
import { FieldValues, SubmitHandler, useForm } from "react-hook-form";
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { BiArrowBack } from "react-icons/bi";
import { Layout } from "../../components/templates/Layout";
import { Input } from "../../components/utils/Input";
import { api } from "../../services/apiClient";

interface FormProps {
    name: string;
    description: string;
    active: boolean;
}

const validation = Yup.object().shape({
    name: Yup.string().required('Nome é obrigatório'),
    description: Yup.string().required('A descrição é obrigatório'),
    //active: Yup.boolean()
    //.required("The terms and conditions must be accepted.s")
    //  .oneOf([true], "The terms and conditions must be accepted.")
})
export function Create() {
    const { register, handleSubmit, reset, formState } = useForm({
        resolver: yupResolver(validation)
    });
    const { errors, isSubmitting } = formState;
    const toast = useToast();
    const navigate = useNavigate();

    const onHandleSubmit: SubmitHandler<FormProps | FieldValues> = async (values) => {
        const formData = new FormData();
        formData.append('name', values.name);
        formData.append('description', values.description);
        formData.append('active', values.active);

        console.log(values)
        const response = await api.post('/campaigns/', {
            name: values.name,
            description: values.description,
            active: values.active
        });
        if (response.status === 201) {
            toast({
                description: "Campanha cadastrada com sucesso!",
                status: 'success',
                duration: 5000,
                isClosable: true
            });
            navigate('/campaigns');
        } else {
            throw new Error("Erro ao salvar informações");

        }
    }


    return (
        <Layout>
            <Flex
                bg="gray.800"
                w="30%"
                p="8"
                mx="auto"
                flexDirection="column"
                borderRadius={8}>
                <Flex
                    onSubmit={handleSubmit(onHandleSubmit)}
                    as="form"
                    flexDirection="column"
                    justify="space-between"
                    h="100%">
                    <Stack spacing={20}>
                        <Flex justify="space-between" align="center">
                            <Heading fontSize="xl">Criar uma campanha</Heading>
                            <Link
                                as={ReactLink}
                                to="/campaigns"
                            >
                                <Icon as={BiArrowBack} fontSize="2xl" />
                            </Link>
                        </Flex>
                        <Stack>
                            <Input
                                {...register('name')}
                                type="text"
                                label="Nome"
                                error={errors.name}
                            />
                            <Input
                                {...register('description')}
                                type="text"
                                label="Descrição"
                                error={errors.description}
                            />
                            <FormControl>
                                <FormLabel>Ativar campanha</FormLabel>
                                <Switch
                                    {...register('active')}
                                />
                                {errors.active && (
                                    <Alert
                                        bg="transparent"
                                        status="error"
                                    >
                                        <AlertIcon />
                                        <AlertDescription>
                                            {/* @ts-ignore */}
                                            {errors.active.message}
                                        </AlertDescription>
                                    </Alert>
                                )}
                            </FormControl>
                        </Stack>
                    </Stack>
                    <Box mt="10">
                        <Button
                            type="submit"
                            transition="filter 0.2s"
                            _hover={{ filter: "brightness(0.9)" }}
                            bg="blue.900"
                        >Salvar
                        </Button>
                    </Box>
                </Flex>


            </Flex>
        </Layout>
    )
}