import { RJSFSchema, UiSchema } from '@rjsf/utils';
import validator from '@rjsf/validator-ajv8';
import Form from '@rjsf/chakra-ui'
import { useEffect, useState } from "react";
import UserLayout from '../layouts/user/user-layout';
import { Container } from '@chakra-ui/react';

interface RegisterFormData {
    email: string,
    password: string,
}

export default function CreateUser() {
    const schema: RJSFSchema = {
        title: "Registro",
        type: 'object',
        required: ['email', 'password'],
        properties: {
            email: {
                type: 'string',
                title: 'Email',
            },
            password: {
                type: 'string',
                title: 'Password',
                
            }
        }
    }

    const uiSchema: UiSchema = {
        password: {
            'ui:widget': 'password',
        }
    }
    return (
        <UserLayout>
            <Container>
                <Form
                    schema={schema}
                    uiSchema={uiSchema}
                    validator={validator}
                    onSubmit={(data, e) => register(data.formData)}
                />
            </Container>
        </UserLayout>
    )
}

function register(data: any) {
    console.log(data)

    fetch('http://localhost:8080/api/user/create', {
        method: "POST",
        body: JSON.stringify(data),
    })
}