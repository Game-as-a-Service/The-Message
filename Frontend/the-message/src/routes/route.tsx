import { createBrowserRouter } from 'react-router-dom';
import Layout from '../layout/layout';
import Map from '../../src/page/game/map';

const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
  },
  {
    path: '/game/map',
    element: <Map />,
  },
]);

export default router;
