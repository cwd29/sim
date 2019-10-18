namespace CoreLib.Compress
{
    interface ICode
    {
        bool Coding(byte[] src,out byte[] dst);
    }
}
