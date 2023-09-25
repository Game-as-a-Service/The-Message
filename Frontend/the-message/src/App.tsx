import { ChakraProvider } from '@chakra-ui/react';
import { RouterProvider } from 'react-router-dom';
import { RecoilRoot } from 'recoil';
import router from '../src/routes/route';

function App() {
  return (
    <RecoilRoot>
      <ChakraProvider>
        <RouterProvider router={router} />
      </ChakraProvider>
    </RecoilRoot>
  );
}

export default App;
