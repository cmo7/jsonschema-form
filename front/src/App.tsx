import { RJSFSchema } from "@rjsf/utils";
import Form from "@rjsf/chakra-ui";
import { useEffect, useState } from "react";
import validator from "@rjsf/validator-ajv8";
import {
  Box,
  Button,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  useDisclosure,
} from "@chakra-ui/react";
//
// const schema: RJSFSchema = {
//   title: "Todo",
//   type: "object",
//   required: ["title"],
//   properties: {
//     title: { type: "string", title: "Title", default: "A new task" },
//     done: { type: "boolean", title: "Done?", default: false },
//     description: { type: "string", title: "Descripción", default: "" },
//   },
// };

interface Todo {
  title: string;
  done: boolean;
  description: string;
}

function App() {
  const [formData, setFormData] = useState({} as Todo);
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [formSchema, setFormSchema] = useState({} as RJSFSchema);
  useEffect(() => {
    (async () => {
      const data = await fetch("http://localhost:8080/get-schema");
      const schema = await data.json();
      console.log(schema);
      setFormSchema(schema);
    })();
  }, []);
  return (
    <>
      <Button onClick={onOpen}>Open Modal</Button>
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Nueva Tarea</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Box p={4}>
              <Form
                schema={formSchema}
                formData={formData}
                validator={validator}
                onChange={(e) => setFormData(e.formData)}
                onError={console.log}
                onSubmit={() => {
                  console.log("submit");
                  onClose();
                }}
              />
            </Box>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
}

export default App;
