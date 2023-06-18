import Form from '@rjsf/chakra-ui';
import { UiSchema } from '@rjsf/utils';
import validator from '@rjsf/validator-ajv8';
import { useState } from 'react';
import { useMutation, useQuery } from 'react-query';
import { endpoints, getSchema, sendForm } from '../api';
import { Loading } from './loading';

export default function Register() {
  const registerSchema = useQuery('schema', async () => getSchema(endpoints.user, 'SiginInput'));
  // TODO: Fix any
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const { mutateAsync } = useMutation(sendForm<any>);
  // TODO: Fix any
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  async function manageSubmit(data: any) {
    const response = await mutateAsync(data);
    console.log(response);
    if (response.status === 'success') {
      window.location.href = '/';
    } else {
      setFormData(data.formData);
    }
  }

  const [formData, setFormData] = useState({});

  const uiSchema: UiSchema = {
    password: {
      'ui:widget': 'password',
    },
    password_confirm: {
      'ui:widget': 'password',
    },
  };
  if (registerSchema.isLoading) return <Loading />;
  return (
    <Form
      schema={registerSchema.data}
      uiSchema={uiSchema}
      validator={validator}
      formData={formData}
      onSubmit={manageSubmit}
    />
  );
}
