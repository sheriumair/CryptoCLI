import './App.css'
import { useState,useEffect,useRef } from 'react';
import axios from 'axios'; 
import { ResponsiveContainer,LineChart,Line } from 'recharts';


export default function App(){
  const [input,setInput]=useState("");
  const [output,setOutput]=useState("");
  const inputRef=useRef();
  const fileInputRef = useRef(null);

  useEffect(()=>{
    inputRef.current.focus() 
  })

  //Here is ABOUT endpoint function
  const handleAboutCommand = async () => {
    try {
      const response = await axios.get('http://localhost:8080/about'); 
      if (response.status === 200) {
        const text = response.data; 
        setOutput(text); 
      } else {
        newOutput('Error: Unable to fetch data from the backend');
      }
    } catch (error) {
      newOutput('Error: Unable to connect to the backend');
    }
  }

    //Here is HELP endpoint function
    const handleHelpCommand = async () => {
      
      try {
        const response = await axios.get('http://localhost:8080/help'); 
        if (response.status === 200) {
          const text = response.data; 
          setOutput(text); 
        } else {
          setOutput('Error: Unable to fetch data from the backend');
        }
      } catch (error) {
        setOutput('Error: Unable to connect to the backend');
      }
    }
        //Here is HELP endpoint function
        const handleFetchCommand = async (pair) => {
            console.log(pair)
          try {
            const response = await axios.get(`http://localhost:8080/fetch-price/${pair}`); 
            if (response.status === 200) {
              const text = response.data; 
              setOutput(text); 
            } else {
              setOutput('Error: Unable to fetch data from the backend');
            }
          } catch (error) {
            setOutput('Error: Unable to connect to the backend');
          }
        }
  
    // Function to fetch data from the backend
    function fetchCommand()  {
        <ResponsiveContainer width="100%" aspect={3}>
          <LineChart>

          </LineChart>
        </ResponsiveContainer>
    };
     //handle UPLOAD
    const handleFileUpload = (e) => {
      console.log("INSIDE FILE")
      const selectedFile = e.target.files[0];
       const formData = new FormData();
      formData.append('file', selectedFile);
  
      axios
        .post('http://localhost:8080/upload', formData)
        .then((response) => {
          setOutput(response.data); // Handle the response from your backend
        })
        .catch((error) => {
          setOutput('Error: Unable to upload file'); // Handle errors
        });
    };


    const handleDrawCommand = async (file, columns) => {
       console.log(file)
       console.log(columns)
    
       try {
        // Make an API request to fetch data from the backend
        const response = await axios.get(`http://localhost:8080/draw?file=${file}&columns=${columns}`);
        
        if (response.status === 200) {
          const data = response.data; // Assuming the backend returns the data in a suitable format
          console.log('Received data from the backend:', data);
    
          // Now you can use 'data' to draw your chart using Recharts
          // Implement your chart drawing logic here
        } else {
          console.error('Error: Unable to fetch data from the backend');
        }
      } catch (error) {
        console.error('Error: Unable to connect to the backend', error);
      }

    };


  return (
    <>
    <div className="App"
      onClick={e=>{inputRef.current.focus()}}
      >
        <h1>Hello WazzzUp</h1>
      <input 
      ref={inputRef}
      type="text"
      value={input}
      onChange={e=>setInput(e.target.value)}
      onKeyDown={async e=>{
        if(e.key === "Enter"){
          let newOutput="";
          newOutput = output + "\n" + "$" + input + "\n"; 
          switch(input){
            case "about":
              //newOutput += "Pinging Backend";
              handleAboutCommand();
              break;
            case "clear":
              ""
              break;
            case "help":
              handleHelpCommand();f
              break;
            /*case "fetch-price":
              console.log("INSIDE")
              const pairInput = input.split(' ')[1]; 
               handleFetchCommand(pairInput);
               break;*/
            case "draw":
              handlefetchCommand();
              break;
            case "upload":
              fileInputRef.current.click()
              break;            
            default:
              if (input.startsWith("fetch-price ")) {
                const pairInput = input.substring(12); // Remove "fetch-price " from the beginning
                console.log(pairInput)
                handleFetchCommand(pairInput);
              } else if (input.startsWith ("draw")) { // Check for the "draw" command
                const parts = input.split(' ');
                if (parts.length >= 3) {
                  const file = parts[1];
                  const columns = parts.slice(2).join(' '); // Join the rest of the parts as columns
                  handleDrawCommand(file, columns);
                } else {
                  newOutput += "Invalid usage. Example: draw [file] [columns]";
                }}
                  else {
                newOutput += "Invalid Command";
              }
          }
          setOutput(newOutput)
          if(input=== "clear"){
            setOutput("")
          }
          setInput("")
        }
      }}
      />
     <input
        ref={fileInputRef}
        type="file"
        style={{ display: "none" }}
        onChange={handleFileUpload}
      />
    <div className="terminal">
    {output}
    </div>
    </div>

    </>

  )
}