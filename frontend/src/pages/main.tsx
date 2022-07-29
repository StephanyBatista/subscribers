import { ChakraProvider } from "@chakra-ui/react";
import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import { AuthProvider } from "../contexts/authContext";

import { AppRoute } from "../services/routes/Router";
import { theme } from "../styles/theme";
import { App } from "./App";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
    <React.StrictMode>
        <BrowserRouter>
            <AuthProvider>
                <ChakraProvider theme={theme}>
                    <AppRoute />
                </ChakraProvider>
            </AuthProvider>
        </BrowserRouter>
    </React.StrictMode>
);
