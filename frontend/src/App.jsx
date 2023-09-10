import './App.css'
import { useState  ,useEffect , useRef } from 'react';
import axios from 'axios'; 
import PopUp  from './components/PopUp'
export default function App() {
  const [input, setInput] = useState("");
  const [output, setOutput] = useState("");
  const inputRef = useRef();
  const [output2, setOutput2] = useState("");
  const [outputArray, setOutputArray] = useState("");
  const [isPopupOpen, setIsPopupOpen] = useState(false);

  const openPopup = () => {
    setIsPopupOpen(true);
  };

  const closePopup = () => {
    setIsPopupOpen(false);
  };
  const fileInputRef = useRef();
  useEffect(() => {
    inputRef.current.focus();
  }, []);
  const handleAboutCommand = async () => {
    try {
      const response = await axios.get("http://localhost:8080/about");
      return response.data;

    } catch (error) {
      return"Error: Unable to connect to the backend";
    }
  };
  const handleHelpCommand = async () => {
    try {
      const response = await axios.get("http://localhost:8080/help");
      return response.data;
    } catch (error) {
      return"Error: Unable to connect to the backend";
    }
  };
  const handleFetchCommand = async (pair) => {
    try {
      const response = await axios.get(
        `http://localhost:8080/fetch-price/${pair}`
      );
      return response.data;
    } catch (error) {
      return error.response.data;
    }
  };
  const handleFileUpload = async (e) => {
    const selectedFile = e.target.files[0];
    const formData = new FormData();
    formData.append("file", selectedFile);
   await axios
      .post("http://localhost:8080/upload", formData)
      .then((response) => {
        setOutput(response.data);
      })
      .catch((error) => {
        setOutput("Error: Unable to upload file");
      });
  };
  const handleDrawCommand = async (file, columns) => {
    try {
      const response = await axios.get(
        `http://localhost:8080/draw?file=${file}&columns=${columns}`
      );
      const data = response.data;
      setOutput2(data)
      setIsPopupOpen(true);
      console.log("Received data from the backend:", data);
      return "Plotted the data";
 
    } catch (error) {
      console.error("Error: Unable to connect to the backend", error);
      return error.response.data
    }
  };

  const handleDeleteCommand = async (fileNameToDelete) => {
    try {
      const response = await axios.delete(`http://localhost:8080/delete/${fileNameToDelete}`);
       return response.data
    } catch (error) {
      return error.response.data;
    }
  };
  const handleInput = async () => {
    let newOutput = output + "\n" + "$ " + input + "\n";
    let response = null;
    switch (input) {
      case "about":
        response = await handleAboutCommand();
        break;
      case "clear":
        newOutput = "";
        break;
      case "help":
        response = await handleHelpCommand();
        break;
      case "upload":
        fileInputRef.current.click();
        break;
      default:
        if (input.startsWith("fetch-price ")) {
          const pairInput = input.substring(12);
          response = await handleFetchCommand(pairInput);
        } else if (input.startsWith("draw")) {
          const parts = input.split(" ");
          if (parts.length >= 3) {
            const file = parts[1];
            const columns = parts.slice(2).join(" ");
            const columnArray = columns.split(',');
            setOutputArray(columnArray)
            console.log(columnArray[0])
            response = await handleDrawCommand(file, columns);
          } else {
            newOutput += "Invalid usage. Example: draw [file] [columns]";
          }
        }  else if (input.startsWith("delete")) {
          const fileName = input.substring(7);
          response = await handleDeleteCommand(fileName);
        } else {
          newOutput += "Invalid Command\n";
        }
    }
    if (response !== null) {
      newOutput += response;
    }
    setOutput(newOutput);
    if (input === "clear") {
      setOutput("");
    }
    setInput("");
  };
  console.log("out" ,output)
  return (
    <>
      <div className="App" onClick={() => inputRef.current.focus()}>
      <div className="HeaderContainer">
        <h1 className="Heading">Welcome to my Terminal.</h1>
        <p className="SubHeading">Write "help" to check the supported commands.</p>
      </div>
        
        <input
          ref={inputRef}
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              handleInput();
            }
          }}
        />
        <input
          ref={fileInputRef}
          type="file"
          style={{ display: "none" }}
          onChange={handleFileUpload}
        />
        <div className="terminal">{output}</div>
      </div>
      <PopUp isOpen={isPopupOpen} onClose={closePopup} output2={output2} outputArray={outputArray}>
        <h2>Popup Content</h2>
        <p>This is the content of the popup.</p>
      </PopUp>

    </>
  );
}