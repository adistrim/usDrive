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
import { useFileUpload } from '@/lib/mutations'

export default function Home() {
    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [isDialogOpen, setIsDialogOpen] = useState(false);
    const fileUploadMutation = useFileUpload();

    const filechangeHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        if (file) {
            setSelectedFile(file);
        }
    }

    const handleUpload = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!selectedFile) return;
        
        fileUploadMutation.mutate(selectedFile, {
            onSuccess: () => {
                setSelectedFile(null);
                setIsDialogOpen(false);
            }
        });
    }

    const handleDialogChange = (open: boolean) => {
        setIsDialogOpen(open);
        if (!open) {
            setSelectedFile(null);
        }
    }

    return (
        <div className="flex justify-end p-5 top-0 items-center pr-16">
            <Dialog open={isDialogOpen} onOpenChange={handleDialogChange}>
                <form>
                    <DialogTrigger asChild>
                        <Button className="cursor-pointer">Upload files</Button>
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
                                disabled={fileUploadMutation.isPending}
                            />
                        </div>
                        <DialogFooter>
                            <Button 
                                type="submit" 
                                disabled={fileUploadMutation.isPending}
                                onClick={handleUpload}
                            >
                                {fileUploadMutation.isPending ? 'Uploading...' : 'Upload'}
                            </Button>
                        </DialogFooter>
                    </DialogContent>
                </form>
            </Dialog>
        </div>
    )
}
