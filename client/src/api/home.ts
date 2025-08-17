import axios from "axios";

export interface FileMetadata {
    fileName: string,
    mimeType: string,
    sizeBytes: number,
    parentId: null,
}

export const requestUploadURL = async (metadata: FileMetadata): Promise<{ uploadUrl: string; fileId: string }> => {
    const res = await axios.post("/api/request/upload", metadata, {
        headers: {
            "Content-Type": "application/json",
        },
    });

    if (res.status !== 200) {
        throw new Error("Failed to request upload URL");
    }

    return res.data;
};


