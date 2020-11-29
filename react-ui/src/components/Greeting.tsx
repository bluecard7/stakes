
import React from 'react';

function Greeting() {
    // Good Afternoon or Good Morning based on local time...?
    return (
        <div>Hi {'{user}'}, today is {new Date().toLocaleDateString()}</div>  
    );
}

export default Greeting;