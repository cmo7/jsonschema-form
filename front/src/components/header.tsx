import {
  Button,
  CloseButton,
  Flex,
  HStack,
  IconButton,
  Link,
  Popover,
  PopoverContent,
  PopoverTrigger,
  Spacer,
  VStack,
  chakra,
  useColorMode,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import { useScroll } from 'framer-motion';
import React from 'react';
import { AiFillHome, AiOutlineInbox, AiOutlineMenu } from 'react-icons/ai';
import { BsFillCameraVideoFill } from 'react-icons/bs';
import { FaMoon, FaSun } from 'react-icons/fa';
import { IoIosArrowDown } from 'react-icons/io';

export default function Header() {
  const { toggleColorMode: toggleMode } = useColorMode();
  const text = useColorModeValue('dark', 'light');
  const SwitchIcon = useColorModeValue(FaMoon, FaSun);
  const bg = useColorModeValue('white', 'gray.800');
  const ref = React.useRef(null);
  const [y, setY] = React.useState(0);
  const height = ref.current ? ref.current.getBoundingClientRect() : 0;
  const { scrollY } = useScroll();
  React.useEffect(() => {
    return scrollY.on('change', () => {
      setY(scrollY.get());
    });
  }, [scrollY]);
  const cl = useColorModeValue('gray.800', 'white');
  const mobileNav = useDisclosure();

  const MobileNavContent = (
    <VStack
      pos="absolute"
      top={0}
      left={0}
      right={0}
      display={mobileNav.isOpen ? 'flex' : 'none'}
      flexDirection="column"
      p={2}
      pb={4}
      m={2}
      bg={bg}
      spacing={3}
      rounded="sm"
      shadow="sm"
    >
      <CloseButton aria-label="Close menu" justifySelf="self-start" onClick={mobileNav.onClose} />
      <Button w="full" variant="ghost" leftIcon={<AiFillHome />}>
        Dashboard
      </Button>
      <Button w="full" variant="solid" colorScheme="brand" leftIcon={<AiOutlineInbox />}>
        Inbox
      </Button>
      <Button w="full" variant="ghost" leftIcon={<BsFillCameraVideoFill />}>
        Videos
      </Button>
    </VStack>
  );
  return (
    <React.Fragment>
      <chakra.header
        ref={ref}
        shadow={y > height ? 'sm' : undefined}
        transition="box-shadow 0.2s"
        bg={bg}
        borderTop="6px solid"
        borderTopColor="brand.400"
        w="full"
        overflowY="hidden"
      >
        <chakra.div h="4.5rem" mx="auto" maxW="1200px">
          <Flex w="full" h="full" px="6" alignItems="center" justifyContent="space-between">
            <Flex align="flex-start">
              <Link href="/">
                <HStack>{/*<Logo />*/}</HStack>
              </Link>
            </Flex>
            <Flex>
              <HStack
                spacing="5"
                display={{
                  base: 'none',
                  md: 'flex',
                }}
              >
                <Popover>
                  <PopoverTrigger>
                    <Button
                      bg={bg}
                      color="gray.500"
                      display="inline-flex"
                      alignItems="center"
                      fontSize="md"
                      _hover={{
                        color: cl,
                      }}
                      _focus={{
                        boxShadow: 'none',
                      }}
                      rightIcon={<IoIosArrowDown />}
                    >
                      Features
                    </Button>
                  </PopoverTrigger>
                  <PopoverContent
                    w="100vw"
                    maxW="md"
                    _focus={{
                      boxShadow: 'md',
                    }}
                  ></PopoverContent>
                </Popover>
                <Button
                  bg={bg}
                  color="gray.500"
                  display="inline-flex"
                  alignItems="center"
                  fontSize="md"
                  _hover={{
                    color: cl,
                  }}
                  _focus={{
                    boxShadow: 'none',
                  }}
                >
                  Blog
                </Button>
                <Button
                  bg={bg}
                  color="gray.500"
                  display="inline-flex"
                  alignItems="center"
                  fontSize="md"
                  _hover={{
                    color: cl,
                  }}
                  _focus={{
                    boxShadow: 'none',
                  }}
                >
                  Pricing
                </Button>
              </HStack>
            </Flex>
            <Spacer />
            <Flex justify="flex-end" align="center" color="gray.400">
              <HStack
                spacing="5"
                display={{
                  base: 'none',
                  md: 'flex',
                }}
              >
                <Button
                  colorScheme="brand"
                  variant="ghost"
                  size="sm"
                  onClick={() => {
                    window.location.href = '/login';
                  }}
                >
                  Acceso
                </Button>
                <Button
                  colorScheme="brand"
                  variant="solid"
                  size="sm"
                  onClick={() => {
                    window.location.href = '/register';
                  }}
                >
                  Registro
                </Button>
              </HStack>
              <IconButton
                size="md"
                fontSize="lg"
                aria-label={`Switch to ${text} mode`}
                variant="ghost"
                color="current"
                ml={{
                  base: '0',
                  md: '3',
                }}
                onClick={toggleMode}
                icon={<SwitchIcon />}
              />
              <IconButton
                display={{
                  base: 'flex',
                  md: 'none',
                }}
                aria-label="Open menu"
                fontSize="20px"
                color="gray.800"
                _dark={{
                  color: 'inherit',
                }}
                variant="ghost"
                icon={<AiOutlineMenu />}
                onClick={mobileNav.onOpen}
              />
            </Flex>
          </Flex>
          {MobileNavContent}
        </chakra.div>
      </chakra.header>
    </React.Fragment>
  );
}
