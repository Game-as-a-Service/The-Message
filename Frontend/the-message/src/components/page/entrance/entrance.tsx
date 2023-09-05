
import Button from "@material-ui/core/Button/Button";
import SendIcon from '@mui/icons-material/Send';
import title from "../../../assets/images/title.png";
import TextField from "@mui/material/TextField/TextField";

const Entrance = () => {

    return (
        <div className="absolute h-4/5 w-96 right-10">
            <div className="flex align-item justify-center w-full">
                <img src={title} className="scale-75"/>
            </div>
            <div className="flex items-center justify-center">
                <TextField id="outlined-basic" color="info" label="Outlined" variant="outlined" />
            </div>
            <div className="p-10 flex items-center justify-center">
                <Button variant="outlined" color="primary" endIcon={<SendIcon />} >
                    進入遊戲
                </Button>
            </div>
        </div>
    );
}

export default Entrance;