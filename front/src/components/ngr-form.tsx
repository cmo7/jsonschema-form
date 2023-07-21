import { Form } from '@rjsf/chakra-ui';
import validator from '@rjsf/validator-ajv8';
import { useQuery } from 'react-query';
import { getSchema, getUiSchema } from '../api/form-schema-requests';
import { customWidgets } from '../rjsf-config/widgets';

type NgrFormProps = {
  schemaName: string;
  formData: any;
  setFormData: (data: any) => void;
  onSubmit?: (data: any) => void;
};

/**
 * This component is a wrapper for the React JSON Schema Form library.
 * It fetches the schema and uiSchema from the server and renders the form.
 * @param schemaName The name of the schema to fetch from the server.
 * @param formData The data to initialize the form with.
 * @param setFormData A function to set the form data.
 * @param onSubmit A function to call when the form is submitted.
 * @returns A React component.
 */
export default function NgrForm({ schemaName, formData, setFormData }: NgrFormProps) {
  // Fetch the schema and uiSchema from the server
  const schema = useQuery(schemaName + '_schema', async () => getSchema(schemaName));
  const uiSchema = useQuery(schemaName + '_uiSchema', async () => getUiSchema(schemaName));

  if (schema.isLoading || uiSchema.isLoading) return <div>Loading...</div>;

  return (
    <Form
      schema={schema.data}
      uiSchema={uiSchema.data}
      formData={formData}
      onSubmit={(data) => {
        console.log(data);
      }}
      onChange={(e) => {
        setFormData(e.formData);
      }}
      validator={validator}
      widgets={customWidgets}
    />
  );
}
