import React from 'react';

export default function ImageForm() {
    return (
    <div>
        <h3>Enter image and params here</h3>
        <form
        encType="multipart/form-data"
        action="http://localhost:9000/api/upload"
        method="post"
        >
        
        <label>File: <input type="file" name="myFile" /></label>
        <label>Desired width in columns: <input type="number" name="width" value={80}/></label>
        <label>Desired height in columns: <input type="number" name="height" value={25}/></label>
        <label>Desired ASCII palette (Most to least dense characters works best here): <input type="string" name="palette" value="'$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\|()1{}[]?-_+~<>i!lI;:,\^`."/></label>
        <label>Title of image: <input type="string" name="title" value="title here"/></label>
        <label>Description of image: <input type="string" name="description" value="description here"/></label>
        <label>Submit here: <input type="submit" value="upload" /></label>
        </form>
    </div>
    )
}