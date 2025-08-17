import axios from "axios";
import { API_URL } from "@/config";

export interface FileMetadata {
    fileName: string,
    mimeType: string,
    sizeBytes: number,
    parentId: null,
}

/*
    requestUploadURL send the reqest to the backend API to request the temp R2 upload URL
*/
export const requestUploadURL = async (metadata: FileMetadata): Promise<{ uploadUrl: string; fileId: string }> => {
    const res = await axios.post(`${API_URL}/request/upload`, metadata, {
        headers: {
            "Content-Type": "application/json",
        },
    });

    if (res.status !== 200) {
        throw new Error("Failed to request upload URL");
    }

    return res.data;
};

/*
    uploadFile uploads the file to the object storage (R2) using the temp link returned the server
*/
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


/*
    completeFileUpload notifies the server about a successful file upload for the record
*/
export const completeFileUpload = async (fileId:string): Promise<void> => {
    const res = await axios.post(`${API_URL}/files/${fileId}/complete`);

    if (res.status !== 200) {
        throw new Error("Failed to complete file upload");
    }
};

