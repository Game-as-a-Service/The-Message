import { Input, Button } from '@chakra-ui/react';
import { Link } from 'react-router-dom';


const Entrance = () => {
  return (
    <div className="absolute h-4/5 w-96 right-10">
      <div className="flex align-item justify-center w-full text-5xl text-white">
        <p className="s">風聲</p>
      </div>
      <div className="flex items-center justify-center mt-5">
        <Input placeholder="輸入玩家姓名" errorBorderColor="red" focusBorderColor="gray.500" color="#fff" />
      </div>
      <div className="p-10 flex items-center justify-center">
        <Button className="bg-gray-500" bg="#66666" color="#fff"><Link to="/game/map">Enter The Game</Link></Button>
      </div>
    </div>
  );
}

export default Entrance;
