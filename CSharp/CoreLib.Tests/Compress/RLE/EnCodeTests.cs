using CoreLib.Compress.RLE;
using NUnit.Framework;
using System;

namespace CoreLib.Tests.Compress.RLE
{
    [TestFixture]
    class EnCodeTests
    {
        [Test]
        public void TwoSameDataTest()
        {
            var code = new EnCode();
            var src = new byte[] { 111, 111 };

            byte[] dst;
            code.Coding(src, out dst);

            Assert.AreEqual(3, dst.Length);
            var exceptData = new byte[] { 2, 111, 111 };
            for (var i = 0; i < exceptData.Length; i++)
                Assert.AreEqual(exceptData[i], dst[i]);
        }

        [Test]
        public void TwoDifferentDataTest()
        {
            var code = new EnCode();
            var src = new byte[] { 110, 111 };

            byte[] dst;
            code.Coding(src, out dst);

            Assert.AreEqual(3, dst.Length);
            var exceptData = new byte[] { 2, 110, 111 };
            for (var i = 0; i < exceptData.Length; i++)
                Assert.AreEqual(exceptData[i], dst[i]);
        }

        [Test]
        public void ThreeSameData()
        {
            var code = new EnCode();
            var src = new byte[] { 111, 111, 111 };

            byte[] dst;
            code.Coding(src, out dst);

            Assert.AreEqual(2, dst.Length);
            var exceptData = new Byte[] { 128 + 3, 111 };
            for (var i = 0; i < exceptData.Length; i++)
                Assert.AreEqual(exceptData[i], dst[i]);
        }

        [Test]
        public void SeriesDifferentData()
        {
            var code = new EnCode();
            var src = new byte[] { 111, 111, 123, 122, 111, 122 };

            byte[] dst;
            code.Coding(src, out dst);


            var exceptData = new byte[] { 6, 111, 111, 123, 122, 111, 122 };
            Assert.AreEqual(exceptData.Length, dst.Length);
            for (var i = 0; i < exceptData.Length; i++)
                Assert.AreEqual(exceptData[i], dst[i]);
        }

        [Test]
        public void BlendAndEndWithSame_1()
        {
            var code = new EnCode();
            var src = new byte[] { 111, 111, 120, 111, 111, 111 };

            byte[] dst;
            code.Coding(src, out dst);

            var exceptData = new byte[] { 3, 111, 111, 120, 128 + 3, 111 };
            Assert.AreEqual(exceptData.Length, dst.Length);
            for (var i = 0; i < exceptData.Length; i++)
                Assert.AreEqual(exceptData[i], dst[i]);
        }

        [Test]
        public void BlendAndEndWithSame_2()
        {
            var code = new EnCode();
            var src = new byte[] { 111, 111, 120, 111, 111, 111, 111 };

            byte[] dst;
            code.Coding(src, out dst);

            var exceptData = new byte[] { 3, 111, 111, 120, 128 + 4, 111 };
            Assert.AreEqual(exceptData.Length, dst.Length);
            for (var i = 0; i < exceptData.Length; i++)
                Assert.AreEqual(exceptData[i], dst[i]);
        }

        [Test]
        public void BlendAndEndWithDifferent_1()
        {
            var code = new EnCode();
            var src = new byte[] { 111, 111, 111, 120, 112, 111, 111 };

            byte[] dst;
            code.Coding(src, out dst);

            var exceptData = new byte[] { 128 + 3, 111, 4, 120, 112, 111, 111 };
            Assert.AreEqual(exceptData.Length, dst.Length);
            for (var i = 0; i < exceptData.Length; i++)
                Assert.AreEqual(exceptData[i], dst[i]);
        }

        [Test]
        public void BlendAndEndWithDifferent_2()
        {
            var code = new EnCode();
            var src = new byte[] { 111, 111, 111, 120, 112, 111 };

            byte[] dst;
            code.Coding(src, out dst);

            var exceptData = new byte[] { 128 + 3, 111, 3, 120, 112, 111 };
            Assert.AreEqual(exceptData.Length, dst.Length);
            for (var i = 0; i < exceptData.Length; i++)
                Assert.AreEqual(exceptData[i], dst[i]);
        }
    }
}
