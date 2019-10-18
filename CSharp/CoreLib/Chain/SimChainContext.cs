using System;
using System.Collections.Generic;

namespace CoreLib.Chain
{
    public class SimChainContext
    {
        public const string KEY_EXIST = "SimChainContext: This key does not exist";

        private Dictionary<string, object> m_Dicts = new Dictionary<string, object>();

        public object this[string key]
        {
            get
            {
                if (this.m_Dicts.ContainsKey(key))
                    return this.m_Dicts[key];

                throw new Exception(KEY_EXIST);
            }
            set
            {
                this.m_Dicts[(string)key] = value;
            }
        }

        public bool HasKey(string key)
        {
            return this.m_Dicts.ContainsKey(key);
        }

        public void SetValue(string key, object value)
        {
            this.m_Dicts.Add(key, value);
        }

        public bool TryGetValue<T>(string key, out T value)
        {
            if (this.m_Dicts.ContainsKey(key))
            {
                value = (T)m_Dicts[key];
                return true;
            }

            value = default(T);
            return false;
        }
    }
}
