#pragma once

#include <iostream>
#include <vector>

class SimpleClassCallback;

class SimpleClass
{
public:
    SimpleClass(){};
    std::string hello();
    void helloString(std::vector<std::string> *results);
    void helloBytes(std::vector<char> *results);
    void setCallBack(SimpleClassCallback *scc);
};

class SimpleClassCallback
{
public:
    virtual ~SimpleClassCallback() {}
    virtual void onStart() {};
};
