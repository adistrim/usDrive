import axios from "axios";

export const uploadFile = async (uploadUrl: string, selectedFile: File) => {
    const response = await axios.put(uploadUrl, selectedFile, {
        headers: {
            'Content-Type': selectedFile.type,
        },
    });
    if (response.status !== 200) {
        throw new Error('Failed to upload file');
    }
}
