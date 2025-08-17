import axios from "axios";

export const completeFileUpload = async (fileId:string): Promise<void> => {
    const res = await axios.post(`/api/files/${fileId}/complete`);

    if (res.status !== 200) {
        throw new Error("Failed to complete file upload");
    }
};
