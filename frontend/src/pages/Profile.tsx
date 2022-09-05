import { Button, Flex, Stack, Tab, TabList, TabPanel, TabPanels, Tabs, useToast } from "@chakra-ui/react";
import { Layout } from "../components/templates/Layout";
import { Input } from "../components/utils/Input";
import { FieldValues, SubmitHandler, useForm } from 'react-hook-form';
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { api } from "../services/apiClient";
import { useAuth } from "../hooks/useAuth";

interface FormProps {
    currentPassword: string;
    newPassword: string;
    confirmedNewPassword: string;
}
const validation = Yup.object().shape({
    currentPassword: Yup.string().required('Digite a senha antiga'),
    newPassword: Yup.string().min(8, 'A senha precisa ter ao menos 8 caracteres').required('Digite uma nova senha'),
    confirmedNewPassword: Yup.string().min(8, 'A senha precisa ter ao menos 8 caracteres').oneOf([Yup.ref('newPassword'), null], 'Senha não confere com a digitada')
})
export function Profile() {
    const { register, handleSubmit, reset, formState } = useForm({
        resolver: yupResolver(validation)
    });
    const { user } = useAuth();
    const { errors, isSubmitting } = formState;
    const toast = useToast();
    const onHandleSubmit: SubmitHandler<FormProps | FieldValues> = async (values) => {

        api.patch(`users/changepassword`, {
            oldpassword: values.currentPassword,
            newpassword: values.newPassword
        }).then((response) => {

            toast({
                description: 'Senha alterada com sucesso!',
                status: 'success',
                duration: 3000,
                isClosable: true
            })
        }).catch(error => {
            console.log(error.response.data.errors[0]);
            toast({
                description: 'Senha antiga não confere!',
                status: 'error',
                duration: 3000,
                isClosable: true
            })
        });

    }
    return (
        <Layout>
            <Flex
                w="100vw"
                bg="gray.800"
                px="10"
                py="8"
                maxWidth={1440}
                h="100vh"
                maxH={680}
            >
                <Tabs w="100%" >
                    <TabList >
                        <Tab>Info</Tab>
                        <Tab>Alterar senha</Tab>
                    </TabList>
                    <TabPanels flex="1" h='100%' >
                        <TabPanel >
                            Info
                        </TabPanel>
                        <TabPanel flex="1" h="100%">
                            <Flex
                                flex="1"
                                as="form"
                                onSubmit={handleSubmit(onHandleSubmit)}
                                h="100%"
                                flexDirection="column"
                            >
                                <Stack
                                    mt="8"
                                    spacing={6}

                                >
                                    <Input
                                        {...register('currentPassword')}
                                        type="password"
                                        label="Senha antiga"
                                        error={errors.currentPassword}
                                    />
                                    <Input
                                        {...register('newPassword')}
                                        type="password"
                                        label="Nova senha"
                                        error={errors.newPassword}
                                    />
                                    <Input
                                        {...register('confirmedNewPassword')}
                                        type="password"
                                        label="Confirmar Senha"
                                        error={errors.confirmedNewPassword}
                                    />
                                </Stack>
                                <Flex
                                    py="8"
                                    h="100%"
                                    flex="1"
                                    justify="flex-start"
                                    align="flex-end"
                                >
                                    <Button
                                        type="submit"
                                        bg="blue.900"
                                        _hover={{ filter: "brightness(0.9)" }}
                                        transition="filter 0.2s"
                                    >Salvar</Button>
                                </Flex>
                            </Flex>

                        </TabPanel>
                    </TabPanels>
                </Tabs>
            </Flex>
        </Layout>
    )
}