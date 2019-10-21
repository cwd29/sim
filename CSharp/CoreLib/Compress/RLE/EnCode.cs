using System;
using System.Collections.Generic;
using System.Linq;

namespace CoreLib.Compress.RLE
{
    public class EnCode : ICode
    {
        public const string ERROR_FORMAT = "RLE Encode Error:{0}";

        public bool Coding(byte[] src, out byte[] dst)
        {
            var result = new List<byte>();
            try
            {
                // 所有的全部都看做不一样的 
                var ds = new List<Byte>();
                var index = 0; // src读取的的下标索引

                var sameCount = 1;
                var sameByte = default(byte);

                ds.Add(src[index++]);
                while (index < src.Length)
                {
                    var b = src[index++];
                    if (ds.Count > 0 && ds.Last() == b) // 不同的话
                    {
                        sameCount++;
                        sameByte = b;

                        if (sameCount >= 3 && ds.Count > 1)
                        {
                            ds.RemoveAt(ds.Count() - 1);
                            result.Add((byte)ds.Count);
                            result.AddRange(ds);
                            ds = new List<byte>();
                            ds.Add(sameByte);
                        }
                    }
                    else
                    {
                        if (sameCount >= 3)
                        {
                            result.Add((byte)(128 + sameCount));
                            result.Add(sameByte);
                            ds = new List<byte>();
                            sameByte = b;
                        }
                        else if (sameCount == 2)
                        {
                            ds.Add(sameByte);
                        }
                        ds.Add(b);
                        sameCount = 1;
                    }
                }
                
                if(sameCount >= 3)
                {
                    ds.RemoveAt(ds.Count() - 1);
                    if(ds.Count() > 0)
                    {
                        result.Add((byte)ds.Count);
                        result.AddRange(ds);
                    }

                    result.Add((byte)(128 + sameCount));
                    result.Add(sameByte);
                }
                else
                {
                    if (sameCount == 2)
                        ds.Add(sameByte);

                    if (ds.Count() > 0)
                    {
                        result.Add((byte)ds.Count);
                        result.AddRange(ds);
                    }
                }
                

            }
            catch (Exception ex)
            {
                throw new Exception(string.Format(ERROR_FORMAT, ex.Message));
            }

            dst = result.ToArray();
            return true;
        }
    }
}
