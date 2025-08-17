import { completeFileUpload, requestUploadURL, uploadFile, type FileMetadata } from '@/api/fileUpload';
import { useMutation } from '@tanstack/react-query';
import { toast } from 'sonner';

export function useFileUpload() {
  return useMutation({
    mutationFn: async (file: File) => {
      const metadata: FileMetadata = {
        fileName: file.name,
        mimeType: file.type,
        sizeBytes: file.size,
        parentId: null,
      };
      
      const { uploadUrl, fileId } = await requestUploadURL(metadata);
      
      await uploadFile(uploadUrl, file);
      
      await completeFileUpload(fileId);
      
      return { fileId };
    },
    onSuccess: () => {
      toast.success("File uploaded successfully!");
    },
    onError: () => {
      toast.error(`Failed to upload file`);
    }
  });
}
