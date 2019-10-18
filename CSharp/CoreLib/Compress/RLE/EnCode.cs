using System;
using System.Collections.Generic;

namespace CoreLib.Compress.RLE
{
    class EnCode : ICode
    {
        public const string ERROR_FORMAT = "RLE Encode Error:{0}";

        public bool Coding(byte[] src, out byte[] dst)
        {
            var result = new List<byte>();
            try
            {
                /**
                 * 1. 连续2个，看做是不同的
                 * */
                var index = 0;
                var coutIndex = 0;
                byte b;
                while (index < src.Length )
                {

                    result[coutIndex] += 1;
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
