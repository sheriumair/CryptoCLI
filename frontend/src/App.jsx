import './App.css'
import { useState,useEffect,useRef } from 'react';
import axios from 'axios'; 
import { ResponsiveContainer,LineChart,Line } from 'recharts';

export default function App(){
  const [input,setInput]=useState("");
  const [output,setOutput]=useState("");
  const inputRef=useRef();

  useEffect(()=>{
    inputRef.current.focus() 
  })

  const handleAboutCommand = async () => {
    try {
      const response = await axios.get('/your-backend-endpoint'); // Replace with your actual backend endpoint
      if (response.status === 200) {
        const text = response.data; // Read the text response from Axios
        setOutput(text); // Set the text as the output
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
              await handleAboutCommand();
              break;
            case "clear":
              ""
            case "fetch":
              await handlefetchCommand();
            default:
              newOutput += "Invalid Command"
          }
          setOutput(newOutput)
          if(input=== "clear"){
            setOutput("")
          }
          setInput("")
        }
      }}
      />
    <div className="terminal"></div>
    {output}
    </div>

    </>

  )
}