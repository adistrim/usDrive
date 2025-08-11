import { Button } from '@/components/ui/button'
import { createFileRoute } from '@tanstack/react-router'
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

export const Route = createFileRoute('/')({
    component: RouteComponent,
})

function RouteComponent() {

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
            const res = await fetch("/api/request/upload", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(metadata),
            })
            const { uploadUrl, fileId } = await res.json();

            await fetch(uploadUrl, {
                method: "PUT",
                body: selectedFile,
            })
            const completeRes = await fetch(`/api/files/${fileId}/complete`, {
                method: "POST",
            })

            const completeupload = await completeRes.json()
            alert("File uploaded successfully!");
            setSelectedFile(null);
        }
        catch (error) {
            alert("Error uploading file: " + error);
        } finally {
            setUploading(false);
        }
    }




    return (

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
    )
}
