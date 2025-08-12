import { Button } from '@/components/ui/button'
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {
    Dialog,
    DialogContent,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import { useState } from 'react'
import axios from 'axios'
import { Toaster } from "@/components/ui/sonner"
import { toast } from "sonner"


export default function Home() {

    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [uploading, setUploading] = useState(false);

    const filechangeHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        if (file) {
            setSelectedFile(file);
        }

    }
    const handleupload = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!selectedFile) return;
        setUploading(true);

        try {
            const metadata = {
                fileName: selectedFile.name,
                mimeType: selectedFile.type,
                sizeBytes: selectedFile.size,
                parentId: null,
            }
            const res = await axios.post("/api/request/upload", metadata, {
                headers: {
                    "Content-Type": "application/json",
                },
            })
            const { uploadUrl, fileId } = await res.data;

            await fetch(uploadUrl, {
                method: "PUT",
                body: selectedFile,
            })
            await axios.post(`/api/files/${fileId}/complete`)
            toast.success("File uploaded successfully!");
            setSelectedFile(null);
        }
        catch (error) {
            toast.error("Error uploading file: " + error);
        } finally {
            setUploading(false);
        }
    }




    return (
        <>
            <div className="flex justify-end p-5 top-0 items-center pr-16">
                <Dialog>
                <form >
                    <DialogTrigger asChild>
                        <Button className="cursor-pointer ">Upload files</Button>
                    </DialogTrigger>
                    <DialogContent className="sm:max-w-[425px]">
                        <DialogHeader>
                            <DialogTitle>Upload your files here!</DialogTitle>
                        </DialogHeader>
                        <div className="space-y-4 mt-4">
                            <Label htmlFor="file-upload">File</Label>
                            <Input
                                id="file-upload"
                                type="file"
                                onChange={filechangeHandler}
                                disabled={uploading}
                            />
                        </div>
                        <DialogFooter>
                            <Button type="submit" onClick={handleupload} disabled={uploading}>Upload</Button>
                        </DialogFooter>
                    </DialogContent>
                </form>
            </Dialog>
        </div>
        <Toaster />
        </>
    )
}
