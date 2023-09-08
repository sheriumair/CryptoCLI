import { useState} from 'react';


 export const handleAboutCommand = async () => {
    const [output,setOutput]=useState("");
  
    try {
      const response = await axios.get('http://localhost:8080/about'); // Replace with your actual backend endpoint
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

