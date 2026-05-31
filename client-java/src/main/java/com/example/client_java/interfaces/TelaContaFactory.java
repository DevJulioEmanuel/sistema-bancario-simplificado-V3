package com.example.client_java.interfaces;

import com.example.client_java.model.ContaLogada;
import com.example.client_java.ui.TelaConta;

@FunctionalInterface
public interface TelaContaFactory {

    TelaConta criar(ContaLogada conta);

}
