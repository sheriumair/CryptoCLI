import '../Popup.css';
import { ResponsiveContainer,LineChart,Line, XAxis,YAxis,Tooltip } from 'recharts';
import PropTypes from 'prop-types';

const PopUp = ({ isOpen, onClose, output2,outputArray }) => {
  return isOpen ? (
    <div className="popup">
      <div className="popup-inner">
        <button className="close-button" onClick={onClose}>
          &times;
        </button>
        <ResponsiveContainer width={1000} aspect={2}>
        <LineChart data={output2} width={500} height={300} margin={{top:5,right:30,left:30,bottom:5}}>
            <XAxis dataKey={outputArray[0]}stroke="purple"/>
            <YAxis />
            <Tooltip contentStyle={{backgroundColor:"yellow"}}/>
          <Line dataKey={outputArray[1]} strok="red" activeDot={{r:8}}/>
          <Line dataKey={outputArray[0]} stroke="green"activeDot={{r:8}}/>
        </LineChart>
      </ResponsiveContainer>
      </div>
    </div>
  ) : null;
};
PopUp.propTypes = {
  isOpen: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  output2: PropTypes.any.isRequired,
  outputArray: PropTypes.array.isRequired,
};

export default PopUp;


