import { ChakraProvider } from '@chakra-ui/react'
import {
  RouterProvider,
} from "react-router-dom";
import router from '../src/routes/route';





const App = () => {
  return (
    <>
      <ChakraProvider>
        <RouterProvider router={router} />
      </ChakraProvider>
    </>
  )
}

export default App
